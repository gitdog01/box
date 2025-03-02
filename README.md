# OX 퀴즈 도우미

OpenAI의 GPT-4o 모델을 활용한 OX 퀴즈 지원 프로그램입니다.

## 주요 기능

- OpenAI API를 사용하여 O/X 퀴즈 질문에 응답
- 간단한 콘솔 인터페이스 제공
- Windows 환경에서 실행 가능

## 빌드 방법

### Mac에서 빌드하기

Mac에서는 다음과 같이 Windows용 실행 파일을 빌드할 수 있습니다:

1. `build_windows.sh` 스크립트를 실행합니다:

   ```
   ./build_windows.sh
   ```

2. 정상적으로 빌드되면 `OX퀴즈도우미.exe` 파일이 생성됩니다.
3. 이 파일을 Windows 사용자에게 전송하면 별도의 설치 없이 실행 가능합니다.

## 사용 방법

1. `OX퀴즈도우미.exe` 파일을 실행합니다.
2. 프롬프트에 OpenAI API 키를 입력합니다 (선택적으로 저장 가능).
3. O/X 퀴즈 질문을 입력하면 AI가 "O" 또는 "X"로 응답합니다.
4. 종료하려면 "exit"를 입력합니다.

## 주의사항

- OpenAI API 사용을 위해서는 유효한 API 키가 필요합니다.
- API 키는 config.json 파일에 저장되거나 환경 변수로 설정할 수 있습니다.
- Windows에서 처음 실행 시 보안 경고가 표시될 수 있으며, "실행" 버튼을 클릭하여 진행할 수 있습니다.

## 필요 조건

- Go 언어 (1.16 이상) - 빌드 시에만 필요
- OpenAI API 키

## API 키 설정 방법

### 1. config.json 파일 사용 (권장)

`config.json` 파일에 API 키를 저장하여 매번 입력하지 않고 사용할 수 있습니다:

```json
{
  "api_key": "여기에_실제_API_키_입력"
}
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

## 예시

```
config.json 파일에서 API 키를 불러왔습니다.
O/X 퀴즈 도우미 프로그램입니다! (종료하려면 'exit' 입력)
질문을 입력하시면 O 또는 X로 답변해 드립니다.

질문을 입력하세요: 대한민국의 수도는 서울인가?
답변: O

질문을 입력하세요: 지구는 평평한가?
답변: X
```
