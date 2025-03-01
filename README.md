# OX 퀴즈 도우미

OpenAI의 GPT-4o 모델을 활용한 OX 퀴즈 지원 프로그램입니다.

## 주요 기능

- OpenAI API를 사용하여 O/X 퀴즈 질문에 응답
- 콘솔 및 그래픽 사용자 인터페이스(GUI) 모드 지원
- Windows 환경에서 실행 가능

## 빌드 방법

### Mac에서 빌드하기

Mac에서는 다음과 같이 빌드할 수 있습니다:

1. `build_windows.sh` 스크립트를 실행합니다:

   ```
   ./build_windows.sh
   ```

2. 스크립트 실행 후 다음 중 선택할 수 있습니다:
   - 옵션 1: 콘솔 모드 애플리케이션 빌드 (GUI 없음)
   - 옵션 2: Windows에서 빌드하기 위한 소스 파일 압축

### Windows에서 빌드하기

Windows에서는 GUI 모드를 포함한 완전한 애플리케이션을 빌드할 수 있습니다:

1. Go 언어를 설치합니다: https://golang.org/dl/
2. 명령 프롬프트(cmd)를 관리자 권한으로 실행합니다.
3. 프로젝트 폴더로 이동합니다.
4. `build_standalone.bat` 파일을 실행합니다:
   ```
   build_standalone.bat
   ```

## 사용 방법

### 콘솔 모드

1. `OX퀴즈도우미_콘솔.exe` 파일을 실행합니다.
2. 프롬프트에 OpenAI API 키를 입력합니다 (선택적으로 저장 가능).
3. O/X 퀴즈 질문을 입력하면 AI가 "O" 또는 "X"로 응답합니다.
4. 종료하려면 "exit"를 입력합니다.

### GUI 모드

1. `OX퀴즈도우미.exe` 파일을 실행합니다.
2. GUI 창에 OpenAI API 키를 입력합니다.
3. 질문 입력 필드에 O/X 퀴즈 질문을 입력합니다.
4. "질문하기" 버튼을 클릭하면 결과가 표시됩니다.

## 주의사항

- OpenAI API 사용을 위해서는 유효한 API 키가 필요합니다.
- API 키는 config.json 파일에 저장되거나 환경 변수로 설정할 수 있습니다.
- Windows에서 처음 실행 시 보안 경고가 표시될 수 있으며, "실행" 버튼을 클릭하여 진행할 수 있습니다.

## 필요 조건

- Go 언어 (1.16 이상)
- OpenAI API 키
- Windows 환경 (GUI 버전)

## 설치 방법

1. 이 저장소를 클론합니다:

   ```
   git clone https://github.com/user/ox-quiz
   cd ox-quiz
   ```

2. 필요한 패키지를 설치합니다:
   ```
   go mod tidy
   ```

## API 키 설정 방법

### 1. config.json 파일 사용 (권장)

`config.json` 파일에 API 키를 저장하여 매번 입력하지 않고 사용할 수 있습니다:

**직접 편집하기**

```json
{
  "api_key": "여기에_실제_API_키_입력"
}
```

**유틸리티 사용하기**

```
set_config.bat
```

### 2. 환경 변수 설정

실행 전 OpenAI API 키를 환경 변수로 설정할 수 있습니다:

**Windows (CMD)**

```
set OPENAI_API_KEY=your-api-key-here
```

**Windows (PowerShell)**

```
$env:OPENAI_API_KEY="your-api-key-here"
```

**영구 환경 변수 설정 (Windows)**

```
set_api_key.bat
```

## 실행 방법

### CLI 버전

#### 직접 실행

```
go run main.go
```

#### 바이너리 빌드 및 실행

```
go build -o ox-quiz.exe
./ox-quiz.exe
```

#### 간편 실행 (Windows)

```
run.bat
```

### GUI 버전 (Windows 전용)

#### 빌드 및 실행

```
build_gui.bat
ox-quiz-gui.exe
```

#### 간편 실행

```
run_gui.bat
```

## 사용 방법

### CLI 버전

1. 프로그램을 실행합니다.
2. API 키가 설정되어 있지 않은 경우 입력하고, 원하면 config.json에 저장할 수 있습니다.
3. O/X 퀴즈 질문을 입력합니다.
4. GPT-4o가 질문에 대해 "O" 또는 "X"로 답변합니다.
5. 종료하려면 "exit"를 입력하세요.

### GUI 버전

1. 프로그램을 실행합니다.
2. OpenAI API 키 필드에 API 키를 입력합니다 (이미 config.json에 저장되어 있으면 자동으로 로드됩니다).
3. 질문 입력란에 O/X 퀴즈 질문을 입력합니다.
4. "질문하기" 버튼을 클릭합니다.
5. 결과가 화면에 표시됩니다.

## 예시

### CLI 버전

```
config.json 파일에서 API 키를 불러왔습니다.
O/X 퀴즈 도우미 프로그램입니다! (종료하려면 'exit' 입력)
질문을 입력하시면 O 또는 X로 답변해 드립니다.

질문을 입력하세요: 대한민국의 수도는 서울인가?
답변: O

질문을 입력하세요: 지구는 평평한가?
답변: X
```

### GUI 버전

![GUI 예시 이미지](gui_example.png)
