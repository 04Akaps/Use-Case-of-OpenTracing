package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

type sample struct {
	n *Network
}

func newSampleRouter(n *Network) {
	s := &sample{n: n}

	n.Router(GET, "/send", s.send)
	n.Router(GET, "/send-with-child", s.sendWithChild)
	n.Router(GET, "/inject", s.inject)
}

func (s *sample) send(c *gin.Context) {
	fmt.Println("Send")

	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan("send_span", ext.RPCServerOption(spanCtx))

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

func (s *sample) inject(c *gin.Context) {
	fmt.Println("inject")

	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan("inject_test", ext.RPCServerOption(spanCtx))

	defer sendSpan.Finish()

	tracer.Inject(sendSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))

}
