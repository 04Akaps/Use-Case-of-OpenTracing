package jager

import (
	"io"
	"open_tracing/jager/provider"
)

type Client struct {
	closer io.Closer
}

func NewClient(service string) *Client {
	c := &Client{}
	var err error

	if c.closer, err = provider.Provider(service); err != nil {
		return nil
	} else {
		return c
	}
}

//
//func (j *Jager) GlobalTracer() opentracing.Tracer {
//	return j.tracer
//}
//
//func (j *Jager) BaseSpan(spanName string) opentracing.Span {
//	return j.GlobalTracer().StartSpan(spanName)
//}
//
//func (j *Jager) BaseCtxSpan(spanName string) context.Context {
//	span := j.BaseSpan(spanName)
//	span.SetTag("api-router", "원하는 API 경로")
//
//	defer span.Finish()
//
//	return opentracing.ContextWithSpan(context.Background(), span)
//}
//
//func (j *Jager) SendSpanUsingTag(spanName string) {
//	span := j.BaseSpan(spanName)
//	span.SetTag("api-router", "원하는 API 경로")
//
//	defer span.Finish()
//
//	//span.LogFields(
//	//	log.String("event", "string-format-log"),
//	//	log.Int("number", 42),
//	//)
//
//	j.UsingChildForUUID(span)
//	//formatString(span, "test")
//}
//
//func (j Jager) SendSpanUsingCtx(spaneName string) {
//	ctx := j.BaseCtxSpan(spaneName)
//
//	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
//	defer span.Finish()
//}
//
//func (j *Jager) UsingChildForUUID(rootSpan opentracing.Span) {
//	// childOf 옵션이 들어가있지 않으면, 16진수 세그먼트의 값이 모두 다르게 된다.
//	// 이러면 로그를 확인하고 추적하는데에 있어서 어려워진다.
//	// 그러니 특수한 상황이 아니라면 RootSpan의 Context를 반드시 넘겨주자.
//	// 해당 형태 말고 ctx를 넘겨주는 형태도 존재한다. -> SendSpanUsingCtx
//
//	span := rootSpan.Tracer().StartSpan(
//		"formatString",
//		opentracing.ChildOf(rootSpan.Context()),
//	)
//
//	defer span.Finish()
//}
