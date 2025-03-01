@echo off
echo OX 퀴즈 도우미 실행 중...

REM 이미 빌드된 실행 파일이 있는지 확인합니다
if exist ox-quiz.exe (
    echo 실행 파일을 실행합니다...
    ox-quiz.exe
) else (
    echo 실행 파일이 없습니다. 빌드 중...
    go build -o ox-quiz.exe
    
    if exist ox-quiz.exe (
        echo 빌드 완료! 실행 중...
        ox-quiz.exe
    ) else (
        echo 빌드 실패! 직접 실행합니다...
        go run main.go
    )
)

pause 