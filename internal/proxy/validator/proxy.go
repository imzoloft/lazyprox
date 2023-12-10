package validator

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/imzoloft/lazyprox/common"
	inout "github.com/imzoloft/lazyprox/internal/io"
	"github.com/imzoloft/lazyprox/internal/io/file"
	"github.com/imzoloft/lazyprox/internal/proxy/validator/helper"
	"github.com/imzoloft/lazyprox/pkg/model"
)

func ValidateProxy() error {
	if err := file.ReadProxies(); err != nil {
		return err
	}

	semaphore := make(chan struct{}, common.Opts.Goroutine)

	var wg sync.WaitGroup

	for _, proxy := range common.Opts.Proxies {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(proxy string) {
			defer func() {
				<-semaphore
				wg.Done()
			}()
			worker(proxy)
		}(proxy)
	}
	wg.Wait()
	close(semaphore)

	return nil
}

func worker(proxy string) {
	if proxy == "" {
		return
	}
	common.Stats.ValidatedProxy++

	proxy = strings.Trim(proxy, "\r")
	parsedProxy, err := ParseProxy(proxy)

	if err != nil {
		fmt.Print(err)
		helper.DeadProxy(proxy)
		return
	}

	if isValidProxy(parsedProxy) {
		helper.WorkingProxy(proxy)
	} else {
		helper.DeadProxy(proxy)
	}
}

func isValidProxy(proxy *model.Proxy) bool {
	resultChan := make(chan bool)
	wg := sync.WaitGroup{}

	for int := 0; int < 2; int++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resultChan <- doRequest(proxy)
		}()
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for int := 0; int < 2; int++ {
		if !<-resultChan {
			return false
		}
	}
	return true
}

func doRequest(proxy *model.Proxy) bool {
	transport, err := GetTransport(proxy)
	if err != nil {
		return false
	}

	if transport == nil {
		return false
	}

	timeoutSeconds := int64(common.Opts.Timeout)
	timeoutDuration := time.Duration(timeoutSeconds) * time.Second
	client := &http.Client{
		Transport: transport,
		Timeout:   timeoutDuration,
	}

	res, err := client.Get("http://ip-api.com/json?fields=8194")

	if err != nil {
		if strings.Contains(err.Error(), "socket: too many open files") {
			inout.FatalError("too many open files (too much goroutines)")
		}
		return false
	}
	defer res.Body.Close()

	return res.StatusCode == 200
}
