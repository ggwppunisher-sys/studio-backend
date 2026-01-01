package apiserver

import (
	"net/http"
	"studio-backend/internal/app/config"
)

func NewServer(cfg config.ServerConfig, handler http.Handler) (*http.Server, error) {
	return &http.Server{
		Addr:           cfg.Port,
		WriteTimeout:   cfg.WriteTimeout,
		ReadTimeout:    cfg.ReadTimeout,
		IdleTimeout:    cfg.IdleTimeout,
		MaxHeaderBytes: cfg.MaxHeaderBytes,

		Handler: handler,
	}, nil
}
