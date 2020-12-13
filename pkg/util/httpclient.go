package util

import (
	"net/http"
	"time"
)

var Client *http.Client

func init() {
	Client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 2,
		},
		Timeout: time.Second * 5,
	}
}
