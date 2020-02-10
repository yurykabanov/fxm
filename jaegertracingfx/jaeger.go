package jaegertracingfx

import (
	"context"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
	"go.uber.org/fx"
)

func JaegerConfigurationProvider() (*jaegerconfig.Configuration, error) {
	defaultConfig := jaegerconfig.Configuration{
		Sampler: &jaegerconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerconfig.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	config, err := defaultConfig.FromEnv()
	if err != nil {
		return nil, err
	}

	return config, nil
}

type TracerCloserHolder struct {
	Closer io.Closer
}

func JaegerTracerProvider(config *jaegerconfig.Configuration) (opentracing.Tracer, *TracerCloserHolder, error) {
	tracer, closer, err := config.NewTracer()
	if err != nil {
		return nil, nil, err
	}

	return tracer, &TracerCloserHolder{Closer: closer}, nil
}

func RegisterTracerAsGlobal(tracer opentracing.Tracer) {
	opentracing.SetGlobalTracer(tracer)
}

func CloseJaegerTracer(lc fx.Lifecycle, closerHolder *TracerCloserHolder) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return closerHolder.Closer.Close()
		},
	})
}
