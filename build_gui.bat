@echo off
echo OX 퀴즈 도우미 GUI 빌드 중...

REM Walk와 rsrc 설치 확인 및 설치
go get github.com/lxn/walk
go get github.com/lxn/win
go get github.com/akavel/rsrc

REM 아이콘과 매니페스트 정보를 담은 .syso 파일 생성
go install github.com/akavel/rsrc@latest
rsrc -manifest ox-quiz.manifest -o rsrc.syso

REM GUI 애플리케이션 빌드
go build -ldflags="-H windowsgui" -o ox-quiz-gui.exe oxgui.go

if exist ox-quiz-gui.exe (
    echo 빌드 완료! ox-quiz-gui.exe 파일이 생성되었습니다.
) else (
    echo 빌드 실패!
)

pause 