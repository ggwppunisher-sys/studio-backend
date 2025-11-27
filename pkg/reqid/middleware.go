package reqid

import (
	"github.com/google/uuid"
	"net/http"
	"strings"
)

const XRequestIDHeader = "X-Request-Id"

func NewHTTPMiddleware(headerName string) func(http.Handler) http.Handler {
	if len(headerName) == 0 {
		headerName = XRequestIDHeader
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := r.Header.Get(headerName)
			if strings.TrimSpace(reqID) == "" {
				reqID = uuid.New().String()
				r.Header.Set(headerName, reqID)
			}
			w.Header().Set(headerName, reqID)
			newContext := ToContext(r.Context(), reqID)
			next.ServeHTTP(w, r.WithContext(newContext))
		})
	}
}
