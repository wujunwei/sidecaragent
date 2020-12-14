package httpserver

import (
	"log"
	"net/http"
	"time"
)

type response struct {
	StatusCode int
	Body       []byte
	header     http.Header
}

// A logger is an http.Handler that logs traffic to standard error.
type logger struct {
	h http.Handler
}
type responseLogger struct {
	code int
	http.ResponseWriter
}

func (r *responseLogger) WriteHeader(code int) {
	r.code = code
	r.ResponseWriter.WriteHeader(code)
}
func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rl := &responseLogger{code: 200, ResponseWriter: w}
	l.h.ServeHTTP(rl, r)
	log.Printf("%.3fs %d %s\n", time.Since(start).Seconds(), rl.code, r.URL)
}

func NewForwardProxy(addr string) *http.Server {
	return &http.Server{Addr: addr, Handler: &logger{h: &ForwardHandle{}}}
}

func NewReverseProxy(addr string) *http.Server {
	return &http.Server{Addr: addr, Handler: &logger{h: &ReverseHandle{}}}
}
