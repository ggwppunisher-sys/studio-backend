package storage

import (
	"database/sql"
	"studio-backend/internal/domain"
)

const (
	createUserQ = ``
	updateUserQ = ``
)

type UsersStorage struct{}

func NewUsersStorage(db *sql.DB) (*UsersStorage, error) {
	return nil, domain.ErrNotImplemented
}
