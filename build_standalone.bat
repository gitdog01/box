@echo off
echo OX 퀴즈 도우미 - 독립 실행형 바이너리 빌드 중...

REM 필요한 패키지 설치
echo 필요한 패키지 설치 중...
go get github.com/lxn/walk
go get github.com/lxn/win
go get github.com/akavel/rsrc

REM 매니페스트 파일로 리소스 생성
echo 리소스 파일 생성 중...
go install github.com/akavel/rsrc@latest
rsrc -manifest ox-quiz.manifest -o rsrc.syso

REM 정적 링크로 단일 바이너리 빌드 (GUI 태그 사용)
echo 독립 실행형 바이너리 빌드 중...
go build -tags "gui walk_use_cgo" -ldflags="-H windowsgui -s -w" -o "OX퀴즈도우미.exe"

REM 임시 파일 정리
del rsrc.syso

REM 빌드 완료 확인
if exist "OX퀴즈도우미.exe" (
    echo 빌드 완료! "OX퀴즈도우미.exe" 파일이 생성되었습니다.
    echo 이 파일은 다른 컴퓨터에서도 별도의 설치 없이 실행 가능합니다.
    echo Windows 보안 경고가 표시될 수 있으니 "실행" 버튼을 클릭하세요.
) else (
    echo 빌드 실패! 오류를 확인하세요.
)

pause 