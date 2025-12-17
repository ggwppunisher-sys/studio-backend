package config

import (
	"studio-backend/internal/domain"
)

type Config struct{}

func New(_, _ string) (Config, error) {
	return Config{}, domain.ErrNotImplemented
}
