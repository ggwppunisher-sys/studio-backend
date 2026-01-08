package handlers

import (
	"context"
	"fmt"
	"studio-backend/internal/domain"
	"studio-backend/internal/transport/apiserver/gen"
	createUser "studio-backend/internal/usecase/create_user"

	"github.com/google/uuid"
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
	ctx context.Context,
	request gen.V1GetUserRequestObject,
) (gen.V1GetUserResponseObject, error) {

	userID := uuid.UUID(request.Id)

	user, err := si.userUseCase.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	apiUser := gen.ObjUser{
		TgId:      &user.TgUserInfo.TgId,
		Username:  ptr(user.TgUserInfo.Username),
		FirstName: ptr(user.TgUserInfo.FirstName),
		LastName:  ptr(user.TgUserInfo.LastName),
		TgChatId:  &user.TgUserInfo.TgChatId,
	}

	return gen.V1GetUser200JSONResponse{
		RespGetUserJSONResponse: gen.RespGetUserJSONResponse{
			Data: apiUser,
		},
	}, nil
}

func (si *StrictImplementation) V1UpdateUser(
	ctx context.Context,
	request gen.V1UpdateUserRequestObject,
) (gen.V1UpdateUserResponseObject, error) {

	userID := uuid.UUID(request.Id)

	user := domain.User{
		Id:         userID,
		TgUserInfo: domain.TgUserInfo{},
	}

	err := si.userUseCase.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	updatedUser, err := si.userUseCase.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	apiUser := gen.ObjUser{
		TgId:      &updatedUser.TgUserInfo.TgId,
		Username:  ptr(updatedUser.TgUserInfo.Username),
		FirstName: ptr(updatedUser.TgUserInfo.FirstName),
		LastName:  ptr(updatedUser.TgUserInfo.LastName),
		TgChatId:  &updatedUser.TgUserInfo.TgChatId,
	}

	return gen.V1UpdateUser200JSONResponse{
		RespUpdateUserJSONResponse: gen.RespUpdateUserJSONResponse{
			Data: apiUser,
		},
	}, nil
}

func ptr(s string) *string {
	return &s
}
