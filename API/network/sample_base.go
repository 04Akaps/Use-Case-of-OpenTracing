package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

func (s *sample) send(c *gin.Context) {
	fmt.Println("Send")

	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan("send_span", ext.RPCServerOption(spanCtx))

	defer sendSpan.Finish()

	c.JSON(http.StatusOK, "Success Sample Span")
}

func (s *sample) sendWithTag(c *gin.Context) {
	fmt.Println("Send With Tag")

	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan("send_span-with-tag",
		opentracing.Tags{
			"tag-test":     "my sample Tag",
			"tag-test two": "my sample Tag Two",
		},
		ext.RPCServerOption(spanCtx),
	)

	defer sendSpan.Finish()

	c.JSON(http.StatusOK, "Success Sample Span")
}

func (s *sample) sendWithChild(c *gin.Context) {
	fmt.Println("sendWithChild")

	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan("span_one", ext.RPCServerOption(spanCtx))

	defer sendSpan.Finish()

	childSpan := tracer.StartSpan("span_two", opentracing.ChildOf(sendSpan.Context()))

	defer childSpan.Finish()
}
