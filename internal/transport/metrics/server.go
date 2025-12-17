package metrics

import (
	"net/http"
	"studio-backend/internal/domain"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewMetricsServer() *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	panic(domain.ErrNotImplemented) // TODO: remove when the server is configured.
	return &http.Server{
		Addr:         "", // TODO
		ReadTimeout:  -1, // TODO
		WriteTimeout: -1, // TODO
		Handler:      mux,
	}
}
