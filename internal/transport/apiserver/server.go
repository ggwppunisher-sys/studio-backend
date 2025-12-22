package apiserver

import (
	"net/http"
	"studio-backend/internal/app/config"
)

func NewServer(handler http.Handler, cfg config.ServerConfig) (*http.Server, error) {
	return &http.Server{
		Handler: handler,

		Addr:         cfg.Port,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}, nil
}
