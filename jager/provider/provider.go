package provider

import (
	"github.com/opentracing/opentracing-go"
	"io"

	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
)

func Provider(service string) (io.Closer, error) {

	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			//LocalAgentHostPort: "localhost:6820", // if host is changed
		},
	}

	if tracer, closer, err := cfg.NewTracer(config.Logger(jaegerLog.StdLogger)); err != nil {
		return nil, err
	} else {
		opentracing.SetGlobalTracer(tracer)

		return closer, err
	}

}
