package httpserver

import "net/http"

var Auth Middleware = func(handler http.Handler) http.Handler {
	return handler
}
