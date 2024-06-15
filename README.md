<h1> Open Tracing </h1>

Golang에서의 OpenTracing 활용법

MSA에서 활용을 하기 떄문에, 상황에 맞춰서 Fx를 활용해
여러가지의 샘플 API 서버를 띄워보는 과정으로 진행 예정



<h1> Tracer 기동 </h1>

현재는 초기 모델이기 때문에 Docker를 활용한다.
- `docker run -d -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest`