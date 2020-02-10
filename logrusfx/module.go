package logrusfx

import (
	"go.uber.org/fx"
)

var (
	DefaultConfigProviderOption = fx.Provide(LoggerConfigProvider)
	DefaultLoggersOption = fx.Provide(LogrusLogger, LogrusFieldLogger, DefaultLoggerAdapter)
)
