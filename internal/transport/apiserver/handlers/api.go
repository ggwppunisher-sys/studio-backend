package handlers

import "studio-backend/internal/domain"

type StrictImplementation struct{}

func NewStrictImplementation() *StrictImplementation {
	panic(domain.ErrNotImplemented)
	return &StrictImplementation{}
}
