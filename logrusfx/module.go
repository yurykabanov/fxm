package logrusfx

import (
	"go.uber.org/fx"
)

var (
	DefaultConfigProvider = fx.Provide(LoggerConfigProvider)
	DefaultLoggers = fx.Provide(LogrusLogger, LogrusFieldLogger, DefaultLoggerAdapter)
)
