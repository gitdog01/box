package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func main() {
	// API 키를 환경 변수에서 가져오거나 직접 입력 받습니다.
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Print("OpenAI API 키를 입력하세요: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		apiKey = scanner.Text()
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