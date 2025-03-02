#!/bin/bash
echo "OX 퀴즈 도우미 - Windows용 실행 파일 빌드 중..."

# 필요한 패키지 설치
echo "필요한 패키지 설치 중..."
go get github.com/sashabaranov/go-openai

# Windows용 크로스 컴파일 실행
echo "Windows용 콘솔 애플리케이션 빌드 중..."
GOOS=windows GOARCH=amd64 go build -o "OX퀴즈도우미.exe" main.go

# 빌드 성공 확인
if [ -f "OX퀴즈도우미.exe" ]; then
  echo "빌드 완료! 'OX퀴즈도우미.exe' 파일이 생성되었습니다."
  echo "이 파일은 Windows 명령 프롬프트 또는 PowerShell에서 실행할 수 있습니다."
  echo "Windows 사용자에게 전송하면 별도의 설치 없이 실행 가능합니다."
else
  echo "빌드 실패! 오류를 확인하세요."
fi 