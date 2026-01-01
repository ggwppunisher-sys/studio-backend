package handlers

import (
	"context"
	"studio-backend/internal/domain"
	"studio-backend/internal/transport/apiserver/gen"
)

type StrictImplementation struct{}

func NewStrictImplementation() *StrictImplementation {
	return &StrictImplementation{}
}

func (si *StrictImplementation) V1CreateUser(
	_ context.Context,
	_ gen.V1CreateUserRequestObject,
) (gen.V1CreateUserResponseObject, error) {
	return nil, domain.ErrNotImplemented
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
