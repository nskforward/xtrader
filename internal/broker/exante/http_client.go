package exante

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var _instance *http.Client

func newClient() *http.Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	return &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			DisableKeepAlives:     false,
			MaxIdleConnsPerHost:   5000,
			MaxConnsPerHost:       5000,
			MaxIdleConns:          5000,
			IdleConnTimeout:       30 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
			TLSHandshakeTimeout:   30 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

func Get() *http.Client {
	if _instance == nil {
		_instance = newClient()
	}
	return _instance
}
