@echo off
echo OpenAI API 키 설정 도구

set /p OPENAI_API_KEY="OpenAI API 키를 입력하세요: "

if "%OPENAI_API_KEY%"=="" (
    echo API 키가 입력되지 않았습니다.
    exit /b
)

setx OPENAI_API_KEY "%OPENAI_API_KEY%"
echo API 키가 시스템 환경 변수로 저장되었습니다.
echo 새 명령 프롬프트 창이나 애플리케이션에서 이 설정이 적용됩니다.

pause 