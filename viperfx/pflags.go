package viperfx

import (
	"os"

	"github.com/spf13/pflag"
)

func PFlagsProvider() (*pflag.FlagSet, error) {
	fs := pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)

	// Config file flag
	fs.StringP("config", "c", "", "Config file")

	err := fs.Parse(os.Args[1:])

	return fs, err
}
