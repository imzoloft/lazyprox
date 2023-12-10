package entry

import (
	"errors"
	"time"

	"github.com/imzoloft/lazyprox/common"
	inout "github.com/imzoloft/lazyprox/internal/io"
	"github.com/imzoloft/lazyprox/pkg/io/file"
	"github.com/imzoloft/lazyprox/pkg/model"
)

func validate(opts *common.Options, timeout int64, proxyType string) error {
	if opts.File == "" {
		return errors.New("no proxy file specified")
	}

	if opts.Output == "" {
		return errors.New("no output file specified")
	}

	if proxyType == "" {
		return errors.New("no proxy type specified")
	}

	if timeout <= 0 {
		return errors.New("invalid timeout value")
	}

	proxyTypeConverted, err := convertProxyType(opts, proxyType)

	if err != nil {
		inout.FatalError(err.Error())
	}
	opts.ProxyType = proxyTypeConverted

	proxies, err := file.ReadFile(opts.File)
	if err != nil {
		return err
	}
	common.Opts.Timeout = time.Duration(timeout) * time.Second
	opts.Proxies = proxies

	return nil
}

func convertProxyType(opts *common.Options, proxyType string) (model.ProxyType, error) {
	switch proxyType {
	case "http":
		return model.HTTP, nil
	case "https":
		return model.HTTPS, nil
	case "socks4":
		return model.SOCKS4, nil
	case "socks5":
		return model.SOCKS5, nil
	default:
		return 0, errors.New("invalid proxy type")
	}
}
