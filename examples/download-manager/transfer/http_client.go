package transfer

import (
	"errors"
	"net"
	"net/http"
	"time"
)

// newClient builds the default client: one dial timeout, one response
// header timeout, and a redirect cap. There is no global request timeout,
// so a large download is not cut off; the stall guard covers a dead socket.
func newClient(opts HTTPOptions) *http.Client {
	connect := opts.ConnectTimeout
	if connect <= 0 {
		connect = 10 * time.Second
	}

	response := opts.ResponseTimeout
	if response <= 0 {
		response = 20 * time.Second
	}

	dialer := &net.Dialer{Timeout: connect, KeepAlive: 30 * time.Second}
	tr := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          4,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   connect,
		ResponseHeaderTimeout: response,
	}

	return &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return errors.New("transfer: too many redirects")
			}

			return nil
		},
	}
}
