package httpserver

import (
	"net/http"
	"sync"
)

var tokenCache sync.Map

type ForwardHandle struct {
}

func (h ForwardHandle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}

type ReverseHandle struct {
}

func (h ReverseHandle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}
