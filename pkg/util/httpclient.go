package util

import (
	"net"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          500,
			IdleConnTimeout:       5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConnsPerHost:   50,
		},
		Timeout: time.Second * 5,
	}
}

func GetAPPHost(appId string, secretToken string) string {
	return ""
}

func GetSecretToken(appId string) string {
	return ""
}
