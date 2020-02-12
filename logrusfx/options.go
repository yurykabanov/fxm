package logrusfx

import (
	"go.uber.org/fx"
)

var (
	DefaultLoggersOption = fx.Provide(LogrusLogger, LogrusFieldLogger, DefaultLoggerAdapter)
)
