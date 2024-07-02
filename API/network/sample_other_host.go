package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

func (s *sample) sendWithOtherHost(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan("root_span", ext.RPCServerOption(spanCtx))
	fmt.Println("send other host")

	fmt.Println("send Header", c.Request.Header)
	fmt.Println()
	defer sendSpan.Finish()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/receive-from-other-host", nil)
	tracer.Inject(sendSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	c.JSON(http.StatusOK, gin.H{"message": "Data sent successfully"})
}

func (s *sample) receiveSpanRouter(c *gin.Context) {
	// router : receive-from-other-host
	tracer := opentracing.GlobalTracer()

	fmt.Println("Receive")
	fmt.Println("receive header", c.Request.Header) // -> Header 정보에 Span의 직렬화된 값 확인 용
	fmt.Println()

	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))

	receiveSpan := tracer.StartSpan("receive_child_span", opentracing.ChildOf(spanCtx))
	defer receiveSpan.Finish()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/receive-two-from-other-host", nil)
	tracer.Inject(receiveSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	c.JSON(http.StatusOK, gin.H{"message": "Span received and processed"})
}

func (s *sample) receiveTwoSpanRouter(c *gin.Context) {
	// router : receive-two-from-other-host
	tracer := opentracing.GlobalTracer()

	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))

	receiveSpan := tracer.StartSpan("receive_child_span_two", opentracing.ChildOf(spanCtx))
	defer receiveSpan.Finish()

	fmt.Println("Receive Two")

	c.JSON(http.StatusOK, gin.H{"message": "Span received and processed"})

}
