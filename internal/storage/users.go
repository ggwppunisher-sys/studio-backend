package storage

import (
	"context"
	"database/sql"
	"errors"
	"studio-backend/internal/domain"

	"github.com/google/uuid"
)

const (
	createUserQ = `
	INSERT INTO um.users (id, first_name, last_name, username, tg_id, tg_chat_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING tg_id 
	`
	getUserQ = `
	SELECT id, first_name, last_name, username, tg_id, tg_chat_id
	FROM um.users
	WHERE id = $1
	`
	updateUserQ = `
	UPDATE um.users
	SET first_name = $2, last_name = $3, username = $4
	WHERE id = $1
	`
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

func (s *UsersStorage) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var u domain.User

	err := s.db.QueryRowContext(ctx, getUserQ, id).Scan(
		&u.Id,
		&u.TgUserInfo.FirstName,
		&u.TgUserInfo.LastName,
		&u.TgUserInfo.Username,
		&u.TgUserInfo.TgId,
		&u.TgUserInfo.TgChatId,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return u, nil
}

func (s *UsersStorage) UpdateUser(ctx context.Context, user domain.User) error {
	res, err := s.db.ExecContext(ctx, updateUserQ,
		user.Id,
		user.TgUserInfo.FirstName,
		user.TgUserInfo.LastName,
		user.TgUserInfo.Username,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user not found for update")

	}
	return nil
}
