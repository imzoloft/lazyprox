package validator

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/imzoloft/lazyprox/common"
	"github.com/imzoloft/lazyprox/internal/proxy/validator/helper"
	"github.com/imzoloft/lazyprox/pkg/model"
	"h12.io/socks"
)

func ParseProxy(proxy string) (*model.Proxy, error) {
	proxyParts := strings.Split(proxy, ":")

	switch len(proxyParts) {
	case 2:
		return &model.Proxy{
			Host: proxyParts[0],
			Port: proxyParts[1],
		}, nil
	case 4:
		return &model.Proxy{
			Host:     proxyParts[0],
			Port:     proxyParts[1],
			Username: proxyParts[2],
			Password: proxyParts[3],
		}, nil
	default:
		return nil, errors.New("invalid proxy format")
	}
}

func GetTransport(proxy *model.Proxy) (*http.Transport, error) {
	transport := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}

	if common.Opts.ProxyType == model.HTTP || common.Opts.ProxyType == model.HTTPS {
		url, err := GetHTTPProxyURL(proxy)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(url)
		return transport, nil
	} else {
		dialer, err := GetSocksDialer(proxy)
		if err != nil {
			return nil, err
		}
		transport.Dial = dialer
		return transport, nil
	}
}

func GetHTTPProxyURL(proxy *model.Proxy) (*url.URL, error) {
	var proxyScheme string

	switch common.Opts.ProxyType {
	case model.HTTP:
		proxyScheme = "http"
	case model.HTTPS:
		proxyScheme = "https"
	default:
		return nil, errors.New("invalid proxy type")
	}

	proxyURL := fmt.Sprintf("%s://%s:%s", proxyScheme, proxy.Host, proxy.Port)
	if helper.IsProxyAuthRequired(proxy) {
		proxyURL = fmt.Sprintf("%s:%s@%s", proxy.Username, proxy.Password, proxyURL)
	}
	return url.Parse(proxyURL)
}

type DialFunc func(string, string) (net.Conn, error)

func GetSocksDialer(proxy *model.Proxy) (DialFunc, error) {
	var proxyURL string

	switch common.Opts.ProxyType {
	case model.SOCKS4:
		proxyURL = fmt.Sprintf("socks4://%s:%s", proxy.Host, proxy.Port)
	case model.SOCKS5:
		if helper.IsProxyAuthRequired(proxy) {
			proxyURL = fmt.Sprintf("socks5://%s:%s@%s:%s", proxy.Username, proxy.Password, proxy.Host, proxy.Port)
		} else {
			proxyURL = fmt.Sprintf("socks5://%s:%s", proxy.Host, proxy.Port)
		}
	default:
		return nil, errors.New("invalid proxy type")
	}

	proxyURL += fmt.Sprintf("?timeout=%s", common.Opts.Timeout.String())
	return socks.Dial(proxyURL), nil
}
