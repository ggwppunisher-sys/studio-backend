package apiserver

import (
	"net/http"
	"studio-backend/internal/domain"
)

func NewServer(handler http.Handler) (*http.Server, error) {
	return nil, domain.ErrNotImplemented
}
