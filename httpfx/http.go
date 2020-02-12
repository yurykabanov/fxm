package httpfx

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type HttpServerConfig struct {
	Addr         string        `mapstructure:"addr"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

type Option func(*HttpServerConfig)

func WithDefault(defaultConfig *HttpServerConfig) Option {
	return func(config *HttpServerConfig) {
		*config = *defaultConfig
	}
}

func WithDefaultAddr(addr string) Option {
	return func(config *HttpServerConfig) {
		config.Addr = addr
	}
}

type HttpServerConfigProviderFunc func(*viper.Viper) (*HttpServerConfig, error)

func MakeHttpServerConfigProvider(rootKey string, opts ...Option) HttpServerConfigProviderFunc {
	return func(v *viper.Viper) (*HttpServerConfig, error) {
		config := &HttpServerConfig{
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		}

		for _, opt := range opts {
			opt(config)
		}

		err := v.UnmarshalKey(rootKey, config)
		if err != nil {
			return nil, err
		}

		if config.Addr == "" {
			return nil, errors.New("empty address is invalid")
		}

		return config, nil
	}
}

func HttpServerProvider(config *HttpServerConfig, logger *log.Logger) *http.Server {
	return &http.Server{
		ErrorLog:     logger,
		Addr:         config.Addr,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
}
