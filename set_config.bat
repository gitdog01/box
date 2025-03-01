@echo off
echo OpenAI API 키 설정 도구 (config.json)

set /p API_KEY="OpenAI API 키를 입력하세요: "

if "%API_KEY%"=="" (
    echo API 키가 입력되지 않았습니다.
    exit /b
)

echo { > config.json
echo     "api_key": "%API_KEY%" >> config.json
echo } >> config.json

echo API 키가 config.json 파일에 저장되었습니다.
echo 이제 프로그램 실행 시 자동으로 이 API 키를 사용합니다.

pause 