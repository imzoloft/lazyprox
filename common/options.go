package common

import (
	"time"

	"github.com/imzoloft/lazyprox/pkg/model"
)

type Options struct {
	Debug     bool
	File      string
	Goroutine int64
	Output    string
	Proxies   []string
	ProxyType model.ProxyType
	StartTime time.Time
	Timeout   time.Duration
	Version   bool
}

type Statistics struct {
	ValidatedProxy int32
	WorkingProxy   int32
	DeadProxy      int32
}
