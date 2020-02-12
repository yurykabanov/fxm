# FX modules

## `viperfx`

Usage:

```go    
import (
    "github.com/yurykabanov/fxm/viperfx"
    "go.uber.org/fx"
)
                       
func main() {
    app := fx.New(
        // ...

        fx.Provide(
            // Enable command line flags:
            // -c /path/to/config/file.ext
            viperfx.PFlagsProvider,

            // Enable default viper config handling
            viperfx.MakeViperProvider(
                // Add prefix to auto env variables
                viperfx.WithEnvPrefix("MY_AWESOME_PREFIX"),

                // Override config name (default is "config.ext")
                viperfx.WithConfigName("my-awesome-config"),

                // Add additional path to search for config
                // Default paths: "." and "./config"
                viperfx.WithAdditionalConfigPaths("./special-config-dir", "/etc/special-config-dir"),

                // Missing config is not an error
                viperfx.WithIgnoreMissingConfig(),
            ),
        )
        
        // ...
    )

    app.Run()
}
```

## `logrusfx`

Depends on `*viper.Viper`.

Usage:

```go    
import (
    "github.com/yurykabanov/fxm/logrusfx"
    "go.uber.org/fx"
)
                       
func main() {
    app := fx.New(
        // ...

        // Enable default config provider
        viperfx.DefaultConfigProviderOption,
                          
        // Enable default loggers: *logrus.Logger, logrus.FieldLogger, log.Logger
        viperfx.DefaultLoggersOption,
        
        // ...
    )

    app.Run()
}
```

## `jaegertracingfx`

You should consider writing your own config provider as the default one may not
suit very well or could require a lot of environment variables.

Usage:

```go      
import (
    "github.com/yurykabanov/fxm/jaegertracingfx"
    "go.uber.org/fx"
)
                       
func main() {
    app := fx.New(
        // ...

        // Enable default config provider
        // It will use very basic config and will parse jaeger's env variables
        jaegertracingfx.DefaultJaegerConfigurationProviderOption,
                          
        // Register Jaeger tracer and a stop-hook
        jaegertracingfx.JaegerTracerOption,
                                                 
        // Register Jaeger tracer as global tracer
        jaegertracingfx.RegisterTracerAsGlobalOption,

        // ...
    )

    app.Run()
}
```

## `httpfx`

Depends on `*viper.Viper`.

```go
import (
    "github.com/yurykabanov/fxm/httpfx"
    "go.uber.org/fx"
)

func main() {
    app := fx.New(
        // ...

        // Enable default config provider
        // It will load config varialbes:
        // http.addr, http.read_timeout, http.write_timeout
        fx.Provide(httpfx.MakeHttpServerConfigProvider("http")),

        // Register `http.Server`
        fx.Provide(httpfx.HttpServerProvider),

        // ...
    )

    app.Run()
}
```
