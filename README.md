<h1> Open Tracing </h1>

Golang에서의 OpenTracing 활용법

MSA에서 활용을 하기 떄문에, 상황에 맞춰서 Fx를 활용해
여러가지의 샘플 API 서버를 띄워보는 과정으로 진행 예정

<h1> 관련된 링크 </h1>

해당 프로젝트와 관련된 나의 공부 일지는 다음과 같다.
- https://impossible-crowley-68b.notion.site/72dd585f46a7422ca962a7bb250a8603?p=8404b7409f184ba48c72b8c785fce05b&pm=c



<h1> Tracer 기동 </h1>

현재는 초기 모델이기 때문에 Docker를 활용한다.
- `docker run -d -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest`
