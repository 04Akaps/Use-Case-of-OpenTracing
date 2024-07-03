package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

func (s *sample) sendForBaggage(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan("root_span", ext.RPCServerOption(spanCtx))

	sendSpan.SetBaggageItem("greeting", "greeting Test")
	fmt.Println("send for baggage")

	defer sendSpan.Finish()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/receive-for-baggage", nil)
	tracer.Inject(sendSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	c.JSON(http.StatusOK, gin.H{"message": "Data sent successfully"})
}

func (s *sample) receiveForBaggage(c *gin.Context) {
	tracer := opentracing.GlobalTracer()

	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))

	receiveSpan := tracer.StartSpan("receive_child_span_for_panic", opentracing.ChildOf(spanCtx))
	defer receiveSpan.Finish()

	fmt.Println(receiveSpan.BaggageItem("greeting"))

	c.JSON(http.StatusOK, gin.H{"message": "Span received and processed"})
}
