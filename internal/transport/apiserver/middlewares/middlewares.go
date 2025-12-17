package middlewares

import (
	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	"net/http"
)

type (
	ClassicMiddleware = func(http.Handler) http.Handler
	StrictMiddleware  = func(next nethttp.StrictHTTPHandlerFunc, operationID string) nethttp.StrictHTTPHandlerFunc
)
