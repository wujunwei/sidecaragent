package httpserver

import "net/http"

type response struct {
	StatusCode int
	Body       []byte
	header     http.Header
}
type Middleware func(handler http.Handler) http.Handler
type AgentHandle struct {
}

func (a AgentHandle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}

func NewServer(addr string) *http.Server {
	return &http.Server{Addr: addr, Handler: &AgentHandle{}}
}
