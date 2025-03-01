//go:build !gui
// +build !gui

package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

// Config 구조체 정의
type Config struct {
	APIKey string `json:"api_key"`
}

// loadConfig 함수는 config.json 파일에서 설정을 읽어옵니다
func loadConfig() (*Config, error) {
	// config.json 파일 읽기
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	// JSON 파싱
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// saveConfig 함수는 설정을 config.json 파일에 저장합니다
func saveConfig(config *Config) error {
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile("config.json", data, 0644)
}

func main() {
	// API 키를 가져오는 순서:
	// 1. config.json 파일에서 읽기 시도
	// 2. 환경 변수에서 읽기 시도
	// 3. 사용자에게 직접 입력 요청

	var apiKey string

	// 1. config.json 파일에서 읽기
	config, err := loadConfig()
	if err == nil && config.APIKey != "" && config.APIKey != "여기에_OpenAI_API_키를_입력하세요" {
		apiKey = config.APIKey
		fmt.Println("config.json 파일에서 API 키를 불러왔습니다.")
	} else {
		// 2. 환경 변수에서 읽기
		apiKey = os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			// 3. 사용자에게 직접 입력 요청
			fmt.Print("OpenAI API 키를 입력하세요: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			apiKey = scanner.Text()

			// 입력받은 API 키를 config 파일에 저장할지 물어봄
			if apiKey != "" {
				fmt.Print("입력한 API 키를 config.json 파일에 저장할까요? (y/n): ")
				scanner.Scan()
				saveChoice := scanner.Text()

				if strings.ToLower(saveChoice) == "y" || strings.ToLower(saveChoice) == "yes" {
					if config == nil {
						config = &Config{}
					}
					config.APIKey = apiKey
					err = saveConfig(config)
					if err != nil {
						fmt.Printf("config.json 파일 저장 중 오류 발생: %v\n", err)
					} else {
						fmt.Println("API 키가 config.json 파일에 저장되었습니다.")
					}
				}
			}
		}
	}

	client := openai.NewClient(apiKey)
	fmt.Println("O/X 퀴즈 도우미 프로그램입니다! (종료하려면 'exit' 입력)")
	fmt.Println("질문을 입력하시면 O 또는 X로 답변해 드립니다.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n질문을 입력하세요: ")
		if !scanner.Scan() {
			break
		}

		userInput := scanner.Text()
		if strings.ToLower(userInput) == "exit" {
			break
		}

		if userInput == "" {
			continue
		}

		// API 요청을 위한 시스템 프롬프트 및 사용자 입력 설정
		answer, err := getOXAnswer(client, userInput)
		if err != nil {
			fmt.Printf("오류 발생: %v\n", err)
			continue
		}

		fmt.Printf("답변: %s\n", answer)
	}
}

func getOXAnswer(client *openai.Client, question string) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4o",
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `당신은 O/X 퀴즈 답변을 제공하는 AI 도우미입니다.
사용자가 O/X 퀴즈 문제를 제시하면 "O" 또는 "X"로만 답변하세요.
정답이 "예"인 경우는 "O"로, 정답이 "아니오"인 경우는 "X"로 답변하세요.
답변은 "O" 또는 "X" 한 글자만 출력하세요.
답변에 대한 이유나 설명은 제공하지 마세요.`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
			MaxTokens: 5, // 짧은 응답만 필요하므로 토큰 수를 제한합니다
		},
	)

	if err != nil {
		return "", err
	}

	// API 응답에서 답변 추출
	answer := resp.Choices[0].Message.Content
	// 공백 및 줄바꿈 제거
	answer = strings.TrimSpace(answer)

	// 답변이 O 또는 X가 아닌 경우 처리
	if answer != "O" && answer != "X" {
		// 답변 내용에 "O"가 포함되어 있으면 O로, "X"가 포함되어 있으면 X로 설정
		if strings.Contains(answer, "O") {
			answer = "O"
		} else if strings.Contains(answer, "X") {
			answer = "X"
		} else {
			// 그 외의 경우 답변을 분석하여 긍정이면 O, 부정이면 X로 설정
			if strings.Contains(strings.ToLower(answer), "yes") ||
				strings.Contains(strings.ToLower(answer), "true") ||
				strings.Contains(strings.ToLower(answer), "맞") {
				answer = "O"
			} else {
				answer = "X"
			}
		}
	}

	return answer, nil
}
