package common

import "time"

const (
	VERSION = "0.1"
	BANNER  = `
 _                 ___  ` + `v` + VERSION + `
| |   __ _ ____  _| _ \_ _ _____ __
| |__/ _` + "`" + ` |_ / || |  _/ '_/ _ \ \ /
|____\__,_/__|\_, |_| |_| \___/_\_\  
              /__/  `
	AUTHOR = `
    gitlab.com/imzoloft
    telegram: @imzoloft`
	USAGE = `
Usage:
  lazyprox -f <proxyfile> -i <outputfile> -p <proxytype>
	
Options:
  FILE
    -f, --file       <proxyfile>    Path to the proxy file
    -o, --output     <outputfile>   Path to the output file
    -p, --proxy      <proxytype>    Type of proxy to use (http, https, socks4, socks5)
   
  HELP
    -h, --help                      Display this help message

  MISC
    -v, --version                   Display the version number		

  SETTINGS
    -d, --debug                     Enable debug mode (show dead proxies)
    -t, --timeout     <timeout>     Timeout in seconds (default: 5)
`
	TextRed   = "\x1b[38;5;124m"
	TextGreen = "\x1b[38;5;42m"
	TextReset = "\x1b[0m"
	TextBlue  = "\x1b[38;5;27m"
)

var Opts = &Options{
	StartTime: time.Now(),
}
var Stats = &Statistics{}
