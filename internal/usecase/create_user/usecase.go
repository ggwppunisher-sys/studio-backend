package createUser

import "studio-backend/internal/domain"

type UseCase struct{}

func New() *UseCase {
	return &UseCase{}
}

func (uc *UseCase) Create() error {
	return domain.ErrNotImplemented
}
