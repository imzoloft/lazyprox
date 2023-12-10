package helper

import (
	"sync/atomic"

	"github.com/imzoloft/lazyprox/common"
	inout "github.com/imzoloft/lazyprox/internal/io"
	"github.com/imzoloft/lazyprox/pkg/io/file"
)

func DeadProxy(proxy string) {
	atomic.AddInt32(&common.Stats.DeadProxy, 1)

	if common.Opts.Debug {
		message := "dead proxy: " + proxy
		inout.DisplayMessage(common.TextRed, message)
	}
}

func WorkingProxy(proxy string) {
	atomic.AddInt32(&common.Stats.WorkingProxy, 1)
	inout.DisplayMessage(common.TextGreen, proxy)
	file.WriteToFile(common.Opts.Output, proxy)
}
