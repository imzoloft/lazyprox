package entry

import (
	"flag"

	"github.com/imzoloft/lazyprox/common"
	inout "github.com/imzoloft/lazyprox/internal/io"
)

func SetupFlag() {
	var (
		proxyType string
		timeout   int64
	)

	flag.BoolVar(&common.Opts.Debug, "d", false, "")
	flag.BoolVar(&common.Opts.Debug, "debug", false, "")

	flag.StringVar(&common.Opts.File, "f", "", "")
	flag.StringVar(&common.Opts.File, "file", "", "")

	flag.Int64Var(&common.Opts.Goroutine, "g", 500, "")
	flag.Int64Var(&common.Opts.Goroutine, "goroutine", 500, "")

	flag.StringVar(&common.Opts.Output, "o", "", "")
	flag.StringVar(&common.Opts.Output, "output", "", "")

	flag.StringVar(&proxyType, "p", "", "")
	flag.StringVar(&proxyType, "proxy", "", "")

	flag.Int64Var(&timeout, "t", 5, "")
	flag.Int64Var(&timeout, "timeout", 5, "")

	flag.BoolVar(&common.Opts.Version, "v", false, "")
	flag.BoolVar(&common.Opts.Version, "version", false, "")

	inout.ClearScreen()
	flag.Usage = func() {
		inout.DisplayBanner()
		inout.DisplayAuthor()
		inout.DisplayUsage()
	}
	flag.Parse()

	if common.Opts.Version {
		inout.DisplayVersion()
	}

	if err := validate(common.Opts, timeout, proxyType); err != nil {
		inout.FatalError(err.Error())
	}
	inout.DisplayBanner()
	inout.DisplayAuthor()
}
