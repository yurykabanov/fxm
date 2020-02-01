package logrusfx

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	ConfigLogLevel  = "log.level"
	ConfigLogFormat = "log.format"
)

type Config struct {
	Level  string
	Format string
}

func LoggerConfigProvider(v *viper.Viper) *Config {
	v.SetDefault(ConfigLogLevel, "info")
	v.SetDefault(ConfigLogFormat, "text")

	return &Config{
		Level:  v.GetString(ConfigLogLevel),
		Format: v.GetString(ConfigLogFormat),
	}
}

func LogrusLogger(config *Config) *logrus.Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		level = logrus.InfoLevel
	}

	logger.SetLevel(level)

	switch config.Format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	default:
		fallthrough
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	return logger
}

func LogrusFieldLogger(logger *logrus.Logger) logrus.FieldLogger {
	return logger
}
