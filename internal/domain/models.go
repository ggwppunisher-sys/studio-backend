package domain

import "github.com/google/uuid"

type TgUserInfo struct {
	TgId      int64
	FirstName string
	LastName  string
	Username  string
	TgChatId  int64
}

func (i TgUserInfo) Validate() error {
	if i.TgId <= 0 {
		return ErrInvalidTgId
	}
	if i.TgChatId <= 0 {
		return ErrInvalidTgChatId
	}
	return nil
}

type User struct {
	Id uuid.UUID

	TgUserInfo
}

func (u User) Validate() error {
	if err := u.TgUserInfo.Validate(); err != nil {
		return err
	}
	if u.Id == uuid.Nil {
		return ErrInvalidUserId
	}
	return nil
}

type Engineer struct {
	User

	EngineerId int64
}

func (e Engineer) Validate() error {
	if err := e.User.Validate(); err != nil {
		return err
	}
	if e.EngineerId <= 0 {
		return ErrInvalidEngineerId
	}
	return nil
}
