package model

type Proxy struct {
	Host     string
	Port     string
	Username string
	Password string
}

type ProxyType int

const (
	HTTP ProxyType = iota + 1
	HTTPS
	SOCKS4
	SOCKS5
)
