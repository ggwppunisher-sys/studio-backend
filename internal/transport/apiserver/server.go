package apiserver

import (
	"net/http"
	"studio-backend/internal/app/config"
	"time"
)

func NewServer(handler http.Handler, cfg config.ServerConfig) (*http.Server, error) {
	return &http.Server{
		Handler: handler,

		Addr:         cfg.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  10 * time.Second,
	}, nil
}
