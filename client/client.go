package troutclient

import (
	"crypto/tls"
	"net/http"

	"golang.org/x/net/proxy"
)

func NewSOCKSDialer(s string) (proxy.ContextDialer, error) {
	dialer, err := proxy.SOCKS5("tcp", s, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}
	return dialer.(proxy.ContextDialer), nil
}

func CreateHTTPClient(s string) (*http.Client, error) {
	if s != "" {
		dialer, err := NewSOCKSDialer(s)
		if err != nil {
			return nil, err
		}
		return &http.Client{
			Transport: &http.Transport{
				DialContext:     dialer.DialContext,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}, nil
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}, nil
}
