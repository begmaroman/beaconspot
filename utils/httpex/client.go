package httpex

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// Default configuration of HTTP client.
const (
	dealTimeout         = time.Minute
	tlsHandshakeTimeout = time.Minute
	clientTimeout       = time.Minute
)

// CreateClientWithTLS creates a new HTTP client with the given TLS configuration.
func CreateClientWithTLS(tls *tls.Config) *http.Client {
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: dealTimeout,
		}).Dial,
		TLSHandshakeTimeout: tlsHandshakeTimeout,
		TLSClientConfig:     tls,
	}

	return &http.Client{
		Timeout:   clientTimeout,
		Transport: netTransport,
	}
}

// CreateClient creates a new HTTP client.
func CreateClient() *http.Client {
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: dealTimeout,
		}).Dial,
		TLSHandshakeTimeout: tlsHandshakeTimeout,
	}

	return &http.Client{
		Timeout:   clientTimeout,
		Transport: netTransport,
	}
}
