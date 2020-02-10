package jaegertracingfx

import (
	"go.uber.org/fx"
)

var (
	// Use default configuration with minimal basic configuration
	// which highly depends on setting correct environment variables
	//
	// See: https://github.com/jaegertracing/jaeger-client-go#environment-variables
	DefaultJaegerConfigurationProviderOption = fx.Provide(JaegerConfigurationProvider)

	// Register Jaeger tracer as `opentracing.Tracer` and stop-hook
	JaegerTracerOption = fx.Options(
		fx.Provide(JaegerTracerProvider),
		fx.Invoke(CloseJaegerTracer),
	)

	// Register `opentracing.Tracer` from container as a global tracer
	RegisterTracerAsGlobalOption = fx.Invoke(RegisterTracerAsGlobal)
)
