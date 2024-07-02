package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type sample struct {
	n *Network
}

func newSampleRouter(n *Network) {
	s := &sample{n: n}

	// sample_base
	n.Router(GET, "/send", s.send)
	n.Router(GET, "/send-with-tag", s.sendWithTag)
	n.Router(GET, "/send-with-child", s.sendWithChild)

	// other_host
	n.Router(GET, "/receive-from-other-host", s.receiveSpanRouter)
	n.Router(GET, "/send-other-host", s.sendWithOtherHost)
	n.Router(GET, "/receive-two-from-other-host", s.receiveTwoSpanRouter)

	n.Router(GET, "/inject", s.inject)
}

func (s *sample) inject(c *gin.Context) {
	fmt.Println("inject")

	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan("inject_test", ext.RPCServerOption(spanCtx))

	defer sendSpan.Finish()

	tracer.Inject(sendSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))

}
