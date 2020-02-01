package viperfx

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type options struct {
	EnvPrefix          string
	DefaultConfigFile  string
	DefaultConfigPaths []string

	IgnoreMissingConfig bool
}

type Option func(opts *options)

func WithEnvPrefix(envPrefix string) Option {
	return func(opts *options) {
		opts.EnvPrefix = envPrefix
	}
}

func WithConfigName(name string) Option {
	return func(opts *options) {
		opts.DefaultConfigFile = name
	}
}

func WithAdditionalConfigPaths(paths ...string) Option {
	return func(opts *options) {
		opts.DefaultConfigPaths = append(opts.DefaultConfigPaths, paths...)
	}
}

func WithIgnoreMissingConfig() Option {
	return func(opts *options) {
		opts.IgnoreMissingConfig = true
	}
}

type ViperProvider func(flagSet *pflag.FlagSet) (*viper.Viper, error)

func MakeViperProvider(opts ...Option) ViperProvider {
	options := options{
		EnvPrefix:          "",
		DefaultConfigFile:  "config",
		DefaultConfigPaths: []string{
			".",
			"./config",
		},
	}

	return func(flagSet *pflag.FlagSet) (*viper.Viper, error) {
		v := viper.New()
		err := v.BindPFlags(flagSet)
		if err != nil {
			return nil, err
		}

		v.AutomaticEnv()

		if options.EnvPrefix != "" {
			v.SetEnvPrefix(options.EnvPrefix)
		}

		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if configFile := v.GetString("config"); configFile != "" {
			// If user do specify config file, then we'll use just this file and won't try default ones

			v.SetConfigFile(configFile)
		} else {
			// If user does not specify config file, then we'll still try to find appropriate config
			// in different folders

			v.SetConfigName(options.DefaultConfigFile)

			for _, dir := range options.DefaultConfigPaths {
				v.AddConfigPath(dir)
			}
		}

		if err := v.ReadInConfig(); err != nil && !options.IgnoreMissingConfig {
			return nil, err
		}

		return v, nil
	}

}
