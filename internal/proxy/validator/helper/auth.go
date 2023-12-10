package helper

import "github.com/imzoloft/lazyprox/pkg/model"

func IsProxyAuthRequired(proxy *model.Proxy) bool {
	return proxy.Username != "" && proxy.Password != ""
}
