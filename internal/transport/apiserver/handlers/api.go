package handlers

import (
	"context"
	"fmt"
	"studio-backend/internal/domain"
	"studio-backend/internal/transport/apiserver/gen"
	createUser "studio-backend/internal/usecase/create_user"
)

type StrictImplementation struct {
	userUseCase *createUser.UseCase
}

func NewStrictImplementation(u *createUser.UseCase) *StrictImplementation {
	return &StrictImplementation{
		userUseCase: u,
	}
}

func (si *StrictImplementation) V1CreateUser(
	ctx context.Context,
	request gen.V1CreateUserRequestObject,
) (gen.V1CreateUserResponseObject, error) {
	body := request.Body
	if body == nil {
		return nil, fmt.Errorf("request body is empty")
	}
	user := domain.User{
		TgUserInfo: domain.TgUserInfo{
			TgId:      body.TgId,
			Username:  valOrEmpty(body.Username),
			FirstName: valOrEmpty(body.FirstName),
			LastName:  valOrEmpty(body.LastName),
			TgChatId:  body.TgChatId,
		},
	}
	createdID, err := si.userUseCase.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	resp := gen.RespCreateUserJSONResponse{}

	resp.Data.UserId = createdID

	return gen.V1CreateUser200JSONResponse{
		RespCreateUserJSONResponse: resp,
	}, nil
}

func valOrEmpty(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func (si *StrictImplementation) V1GetUser(
	_ context.Context,
	_ gen.V1GetUserRequestObject,
) (gen.V1GetUserResponseObject, error) {

	return nil, domain.ErrNotImplemented
}

func (si *StrictImplementation) V1UpdateUser(
	_ context.Context,
	_ gen.V1UpdateUserRequestObject,
) (gen.V1UpdateUserResponseObject, error) {
	return nil, domain.ErrNotImplemented
}
