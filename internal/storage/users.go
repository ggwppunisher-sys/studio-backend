package storage

import (
	"context"
	"database/sql"
	"studio-backend/internal/domain"
)

const (
	createUserQ = `
	INSERT INTO um.users (id, first_name, last_name, username, tg_id, tg_chat_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING tg_id 
	`
	updateUserQ = ``
)

type UsersStorage struct {
	db *sql.DB
}

func NewUsersStorage(db *sql.DB) (*UsersStorage, error) {
	return &UsersStorage{db: db}, nil
}

func (s *UsersStorage) SaveUser(ctx context.Context, user domain.User) (int64, error) {
	var id int64
	err := s.db.QueryRowContext(ctx, createUserQ,
		user.Id,
		user.TgUserInfo.FirstName,
		user.TgUserInfo.LastName,
		user.TgUserInfo.Username,
		user.TgUserInfo.TgId,
		user.TgUserInfo.TgChatId,
	).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}
