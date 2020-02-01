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

Usage:

```go    
import (
    "github.com/yurykabanov/fxm/logrusfx"
    "go.uber.org/fx"
)
                       
func main() {
    app := fx.New(
        // ...

        fx.Provide(
            // Enable default config provider
            viperfx.DefaultConfigProvider,         
                          
            // Enable default loggers: *logrus.Logger, logrus.FieldLogger, log.Logger
            viperfx.DefaultLoggers,
        )
        
        // ...
    )

    app.Run()
}
```
