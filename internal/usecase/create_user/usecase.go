package createUser

import (
	"context"
	"studio-backend/internal/domain"

	"github.com/google/uuid"
)

type UserSaver interface {
	SaveUser(ctx context.Context, user domain.User) (int64, error)
}
type UseCase struct {
	saver UserSaver
}

func New(saver UserSaver) *UseCase {
	return &UseCase{
		saver: saver,
	}
}

func (uc *UseCase) Create(ctx context.Context, user domain.User) (uuid.UUID, error) {
	user.Id = uuid.New()
	if err := user.Validate(); err != nil {
		return uuid.Nil, err
	}
	_, err := uc.saver.SaveUser(ctx, user)
	if err != nil {
		return uuid.Nil, err
	}
	return user.Id, nil
}
