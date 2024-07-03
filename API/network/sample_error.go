package network

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
)

func (s *sample) sendForPanic(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan("root_span", ext.RPCServerOption(spanCtx))
	fmt.Println("send for panic")

	defer sendSpan.Finish()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/receive-for-error", nil)
	tracer.Inject(sendSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	c.JSON(http.StatusOK, gin.H{"message": "Data sent successfully"})
}

func (s *sample) receiveForError(c *gin.Context) {
	// router : receive-for-error
	tracer := opentracing.GlobalTracer()

	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))

	receiveSpan := tracer.StartSpan("receive_child_span_for_panic", opentracing.ChildOf(spanCtx))
	defer receiveSpan.Finish()

	err := errors.New("sample error")

	receiveSpan.SetTag("error", true)
	receiveSpan.LogFields(log.String("event", "error"), log.String("message", err.Error()))

	c.JSON(http.StatusOK, gin.H{"message": "Span received and processed"})
}
