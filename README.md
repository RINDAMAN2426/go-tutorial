# Chapter 1 - Go 시작하기

- multiple 모듈의 경우 뒤 챕터 내용에서 서술됨

- `go mod tidy`를 사용해서, 사용중인 혹은 사용하지 않는 모듈에 대한 관리

- package는 관리중이며 Go 도구들이 다운로드 받을 수 있는 경로로 서술해야함.

# Chapter 2 - 모듈 만들기

- export하기 위한 모듈 생성하기. package명을 `main`을 사용하면 안됨. 그러면 호출 모듈에서 해당 모듈이 프로그램이라고 인식해버림

- export하기 위해선 앞글자를 대문자로 사용해야함. [Exported names](https://go.dev/tour/basics/3) 참고하기

- 타입은 코드대로 정의하면 됨

- `main` 패키지 아니면 실행 안됨. (그러면 main으로 두고 나중에 패키지명 바꾸는 식으로 개발하나..? 테스트 코드 작성 쯤 가면 알듯)

# Chapter 3 - 모듈 호출하기

- 해당 모듈이 publish된 경우는 해당 URL을 사용하면 되지만, 로컬 개발 시에는 publish 되지 않았기 때문에 replace로 대처함. 버저닝은 맞춰주면 됨

