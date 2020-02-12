package httpfx

import (
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type HttpServerConfig struct {
	Addr         string        `mapstructure:"addr"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

type HttpServerConfigProviderFunc func(v *viper.Viper) (*HttpServerConfig, error)

func MakeHttpServerConfigProvider(rootKey string) HttpServerConfigProviderFunc {
	return func(v *viper.Viper) (*HttpServerConfig, error) {
		config := HttpServerConfig{}

		err := v.UnmarshalKey(rootKey, &config)
		if err != nil {
			return nil, err
		}

		return &config, nil
	}
}

func HttpServerProvider(config *HttpServerConfig) *http.Server {
	return &http.Server{
		Addr:         config.Addr,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
}
