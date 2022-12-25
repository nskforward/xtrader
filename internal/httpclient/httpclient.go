package httpclient

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func InitDefault() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Jar = jar
	http.DefaultClient.Transport = &http.Transport{
		DisableKeepAlives:     false,
		MaxIdleConnsPerHost:   10,
		MaxConnsPerHost:       100,
		MaxIdleConns:          50,
		IdleConnTimeout:       2 * time.Minute,
		ResponseHeaderTimeout: 10 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
}
