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
		},
	}

	if tracer, closer, err := cfg.NewTracer(config.Logger(jaegerLog.StdLogger)); err != nil {
		return nil, err
	} else {
		opentracing.SetGlobalTracer(tracer)

		return closer, err
	}

	//tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	//if err != nil {
	//	panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	//}
	//return tracer, closer
	//
	//client := otlptracegrpc.NewClient(
	//	otlptracegrpc.WithInsecure(),
	//)
	//exporter, err := otlptrace.New(ctx, client)
	//if err != nil {
	//	log.Fatal("creating OTLP trace exporter: %w", err)
	//}
	//
	//tp := sdk.NewTracerProvider(
	//	sdk.WithBatcher(exporter),
	//	sdk.WithResource(newResource(service)),
	//)
	//
	//return tp.Tracer(service)
}

//func newResource(service string) *resource.Resource {
//	return resource.NewWithAttributes(
//		conv.SchemaURL,
//		conv.ServiceName(service),
//		conv.ServiceVersion("0.0.1"),
//	)
//}
