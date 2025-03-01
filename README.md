# OX 퀴즈 도우미

OpenAI GPT-4o를 이용한 O/X 퀴즈 답변 생성 프로그램입니다.

## 필요 조건

- Go 언어 (1.16 이상)
- OpenAI API 키

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

## 실행 방법

### 직접 실행

```
go run main.go
```

### 바이너리 빌드 및 실행

```
go build -o ox-quiz.exe
./ox-quiz.exe
```

### 환경 변수 설정

실행 전 OpenAI API 키를 환경 변수로 설정할 수 있습니다:

**Windows (CMD)**

```
set OPENAI_API_KEY=your-api-key-here
```

**Windows (PowerShell)**

```
$env:OPENAI_API_KEY="your-api-key-here"
```

## 사용 방법

1. 프로그램을 실행합니다.
2. OpenAI API 키를 입력하거나 환경 변수로 설정합니다.
3. O/X 퀴즈 질문을 입력합니다.
4. GPT-4o가 질문에 대해 "O" 또는 "X"로 답변합니다.
5. 종료하려면 "exit"를 입력하세요.

## 예시

```
O/X 퀴즈 도우미 프로그램입니다! (종료하려면 'exit' 입력)
질문을 입력하시면 O 또는 X로 답변해 드립니다.

질문을 입력하세요: 대한민국의 수도는 서울인가?
답변: O

질문을 입력하세요: 지구는 평평한가?
답변: X
```
