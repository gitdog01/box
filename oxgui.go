//go:build gui && windows
// +build gui,windows

// 독립적인 GUI 애플리케이션 패키지
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/sashabaranov/go-openai"
)

// GUIConfig 구조체 정의
type GUIConfig struct {
	APIKey string `json:"api_key"`
}

// 애플리케이션 상태 저장을 위한 구조체
type OXQuizApp struct {
	window            *walk.MainWindow
	questionInput     *walk.TextEdit
	resultOutput      *walk.TextEdit
	submitButton      *walk.PushButton
	apiKeyInput       *walk.LineEdit
	statusLabel       *walk.Label
	client            *openai.Client
	apiKey            string
	apiKeyLastChecked bool
}

// loadGUIConfig 함수는 config.json 파일에서 설정을 읽어옵니다
func loadGUIConfig() (*GUIConfig, error) {
	// config.json 파일 읽기
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	// JSON 파싱
	var config GUIConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// saveGUIConfig 함수는 설정을 config.json 파일에 저장합니다
func saveGUIConfig(config *GUIConfig) error {
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile("config.json", data, 0644)
}

func (app *OXQuizApp) sendQuestion() {
	// API 키 확인
	apiKey := strings.TrimSpace(app.apiKeyInput.Text())

	// API 키가 변경되었으면 클라이언트 업데이트
	if apiKey != app.apiKey {
		app.apiKey = apiKey
		app.client = openai.NewClient(apiKey)

		// API 키 저장
		config := &GUIConfig{APIKey: apiKey}
		err := saveGUIConfig(config)
		if err != nil {
			app.statusLabel.SetText("API 키 저장 중 오류 발생: " + err.Error())
			return
		}
	}

	// API 키가 비어있는지 확인
	if app.apiKey == "" {
		app.statusLabel.SetText("API 키를 입력해주세요.")
		return
	}

	question := app.questionInput.Text()
	if question == "" {
		app.statusLabel.SetText("질문을 입력해주세요.")
		return
	}

	app.statusLabel.SetText("질문을 처리 중입니다...")
	app.submitButton.SetEnabled(false)

	// 고루틴에서 API 호출 실행
	go func() {
		answer, err := getGUIOXAnswer(app.client, question)

		// UI 업데이트는 메인 스레드에서 실행
		// walk.Synchronize 대신 window.Synchronize 메소드 사용
		app.window.Synchronize(func() {
			app.submitButton.SetEnabled(true)

			if err != nil {
				app.statusLabel.SetText("오류 발생: " + err.Error())
				return
			}

			// 결과 표시
			app.resultOutput.SetText(answer)
			app.statusLabel.SetText("응답 완료")
		})
	}()
}

func getGUIOXAnswer(client *openai.Client, question string) (string, error) {
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

// GUI 애플리케이션의 메인 함수
func main() {
	// 애플리케이션 객체 생성
	app := &OXQuizApp{}

	// config.json에서 API 키 로드 시도
	config, err := loadGUIConfig()
	if err == nil && config.APIKey != "" && config.APIKey != "여기에_OpenAI_API_키를_입력하세요" {
		app.apiKey = config.APIKey
	} else {
		// 환경 변수에서 API 키 로드 시도
		app.apiKey = os.Getenv("OPENAI_API_KEY")
	}

	// OpenAI 클라이언트 생성
	if app.apiKey != "" {
		app.client = openai.NewClient(app.apiKey)
	}

	// 메인 윈도우 생성
	if err := (MainWindow{
		AssignTo: &app.window,
		Title:    "OX 퀴즈 도우미",
		MinSize:  Size{Width: 500, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			// API 키 입력 섹션
			GroupBox{
				Title:  "OpenAI API 키",
				Layout: HBox{},
				Children: []Widget{
					LineEdit{
						AssignTo:     &app.apiKeyInput,
						Text:         app.apiKey,
						PasswordMode: true,
						MinSize:      Size{Width: 300},
					},
				},
			},

			// 질문 입력 섹션
			GroupBox{
				Title:  "질문 입력",
				Layout: VBox{},
				Children: []Widget{
					TextEdit{
						AssignTo: &app.questionInput,
						MinSize:  Size{Height: 100},
						Font:     Font{Family: "맑은 고딕", PointSize: 10},
					},
					PushButton{
						AssignTo: &app.submitButton,
						Text:     "질문하기",
						OnClicked: func() {
							app.sendQuestion()
						},
						MinSize: Size{Width: 100},
					},
				},
			},

			// 결과 표시 섹션
			GroupBox{
				Title:  "결과",
				Layout: VBox{},
				Children: []Widget{
					TextEdit{
						AssignTo:      &app.resultOutput,
						ReadOnly:      true,
						MinSize:       Size{Width: 400, Height: 100},
						Font:          Font{Family: "맑은 고딕", PointSize: 24, Bold: true},
						TextAlignment: AlignCenter,
					},
				},
			},

			// 상태 표시줄
			Label{
				AssignTo: &app.statusLabel,
				Text:     "준비됨",
			},
		},
	}.Create()); err != nil {
		fmt.Println("창 생성 오류:", err)
		os.Exit(1)
	}

	// 윈도우 표시
	app.window.Run()
}
