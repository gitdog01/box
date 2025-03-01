#!/bin/bash
echo "OX 퀴즈 도우미 - Windows용 실행 파일 빌드 중..."

# 필요한 패키지 설치
echo "필요한 패키지 설치 중..."
go get github.com/sashabaranov/go-openai

# 매니페스트 파일 확인
if [ ! -f "ox-quiz.manifest" ]; then
  echo "매니페스트 파일 생성 중..."
  cat > ox-quiz.manifest << EOL
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
  <assemblyIdentity
      version="1.0.0.0"
      processorArchitecture="*"
      name="OX퀴즈도우미"
      type="win32"
  />
  <description>OX 퀴즈 도우미</description>
  <trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
    <security>
      <requestedPrivileges>
        <requestedExecutionLevel
            level="asInvoker"
            uiAccess="false"
        />
      </requestedPrivileges>
    </security>
  </trustInfo>
  <dependency>
    <dependentAssembly>
      <assemblyIdentity
          type="win32"
          name="Microsoft.Windows.Common-Controls"
          version="6.0.0.0"
          processorArchitecture="*"
          publicKeyToken="6595b64144ccf1df"
          language="*"
      />
    </dependentAssembly>
  </dependency>
</assembly>
EOL
  echo "매니페스트 파일을 생성했습니다."
fi

echo "중요: Mac에서는 walk 라이브러리를 사용하는 Windows GUI 애플리케이션을 직접 빌드할 수 없습니다."
echo "다음 두 가지 방법 중 하나를 선택하세요:"
echo ""
echo "1. 콘솔 모드 애플리케이션만 빌드하기 (GUI 없음)"
echo "2. Windows 환경에서 build_standalone.bat 파일을 실행하여 빌드하기"
echo ""
read -p "옵션을 선택하세요 (1 또는 2): " option

if [ "$option" = "1" ]; then
  echo "Windows용 콘솔 애플리케이션 빌드 중..."
  # 콘솔 애플리케이션 빌드
  GOOS=windows GOARCH=amd64 go build \
    -tags "!gui" \
    -o "OX퀴즈도우미_콘솔.exe" \
    main.go

  # 빌드 성공 확인
  if [ -f "OX퀴즈도우미_콘솔.exe" ]; then
    echo "빌드 완료! 'OX퀴즈도우미_콘솔.exe' 파일이 생성되었습니다."
    echo "이 파일은 Windows 명령 프롬프트 또는 PowerShell에서 실행할 수 있습니다."
  else
    echo "빌드 실패! 오류를 확인하세요."
  fi
else
  echo ""
  echo "Windows에서 build_standalone.bat 파일을 실행하여 GUI 애플리케이션을 빌드하세요."
  echo "필요한 파일들을 Windows 시스템으로 전송한 후 다음 절차를 따르세요:"
  echo ""
  echo "1. Go 언어 설치: https://golang.org/dl/"
  echo "2. 명령 프롬프트(cmd)를 관리자 권한으로 실행"
  echo "3. 프로젝트 폴더로 이동"
  echo "4. build_standalone.bat 실행"
  echo ""
  echo "모든 파일을 압축하여 Windows 사용자에게 전송하시겠습니까? (y/n)"
  read -p "> " create_zip
  
  if [ "$create_zip" = "y" ] || [ "$create_zip" = "Y" ]; then
    # ZIP 파일 생성
    echo "프로젝트 파일 압축 중..."
    zip -r OX퀴즈도우미_소스.zip . -x "*.git*" "*.exe" "*.zip"
    echo "OX퀴즈도우미_소스.zip 파일이 생성되었습니다."
    echo "이 파일을 Windows 사용자에게 전송한 후 압축을 풀고 build_standalone.bat를 실행하도록 안내하세요."
  fi
fi 