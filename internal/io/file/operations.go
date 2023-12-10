package file

import (
	"github.com/imzoloft/lazyprox/common"
	"github.com/imzoloft/lazyprox/pkg/io/file"
)

func ReadProxies() error {
	proxies, err := file.ReadFile(common.Opts.File)
	if err != nil {
		return err
	}
	common.Opts.Proxies = proxies
	return nil
}
