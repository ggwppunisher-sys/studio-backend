package startup

import (
	"context"
	"studio-backend/internal/app/config"
	"studio-backend/internal/domain"
)

func Run(_ context.Context, _ config.Config) error {
	return domain.ErrNotImplemented
}
