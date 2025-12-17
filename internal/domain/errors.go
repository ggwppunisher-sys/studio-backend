package domain

import (
	"fmt"
)

type Error struct {
	inner        error
	businessCode string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.inner.Error(), e.businessCode)
}

func (e Error) Unwrap() error {
	return e.inner
}

func (e Error) Code() string {
	return e.businessCode
}

type errSemantic string

func (e errSemantic) Error() string {
	return string(e)
}

func (e errSemantic) withCode(code string) Error {
	return Error{
		inner:        e,
		businessCode: code,
	}
}

const (
	ErrInvalidData  errSemantic = "invalid_data"
	ErrConflict     errSemantic = "conflict"
	ErrBadPayload   errSemantic = "bad_payload"
	ErrUnauthorized errSemantic = "unauthorized"
	ErrForbidden    errSemantic = "forbidden"
	ErrNotFound     errSemantic = "not_found"
	ErrExpired      errSemantic = "expired"
	ErrUnsupported  errSemantic = "unsupported"
	ErrInternal     errSemantic = "internal_error"
)

var (
	ErrNotImplemented = ErrInternal.withCode("not_implemented")
)
