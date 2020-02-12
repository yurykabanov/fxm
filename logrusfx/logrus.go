package logrusfx

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

type LoggerConfigProviderFunc func (v *viper.Viper) (*Config, error)

func MakeLoggerConfigProvider(rootKey string) LoggerConfigProviderFunc {
	return func (v *viper.Viper) (*Config, error) {
		config := &Config{
			Level:  "info",
			Format: "text",
		}

		err := v.UnmarshalKey(rootKey, config)
		if err != nil {
			return nil, err
		}

		return config, nil
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
