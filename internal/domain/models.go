package domain

import (
	"errors"
	"time"
)

type User struct {
	Id int64

	TgUserInfo
}

type TgUserInfo struct {
	TgId      int64
	FirstName string
	LastName  string
	Username  string
	TgChatId  int64
}

type Engineer struct {
	User
}

type Schedule struct {
	Start    time.Time // Start time of recording
	Finish   time.Time // Finish time of recording
	Engineer Engineer  // Which Engineer
	FreeTime time.Time // Free time for recording
}

type Reservation struct {
	PlaceId      string    // Place where will be recording
	Time         time.Time // When
	UserTgId     string    // User telegram UserName
	EngineerTgId string    // Engineer telegram UserName
	Duration     time.Time // Duration of recording
}

type BookingScheduleEngineer struct {
	TgId    string    // telegram UserName
	PlaceId string    // Place where will be recording
	Time    time.Time // Duration of recording

}

type BookingScheduleUser struct {
	Place      string    // Where will be recording
	Time       time.Time // Duration of recording
	EngineerId string    // Engineer ID
}

func (u *User) Validate() error {

	return nil
}

func (e *Engineer) Validate() error {

	return nil
}

func (r *Reservation) Validate() error {
	if r.UserTgId == "" {
		return errors.New("Incorrectly entered field")
	}
	if r.EngineerTgId == "" {
		return errors.New("Incorrectly entered field")
	}
	return nil
}

func (b *BookingScheduleEngineer) Validate() error {
	if b.TgId == "" {
		return errors.New("Incorrectly entered field")
	}
	if b.PlaceId == "" {
		return errors.New("Incorrectly entered field")
	}
	return nil
}

func (bo *BookingScheduleUser) Validate() error {
	if bo.Place == "" {
		return errors.New("Incorrectly entered field")
	}
	if bo.EngineerId == "" {
		return errors.New("Incorrectly entered field")
	}

	return nil
}

func (s *Schedule) Validate() error {
	return nil
}
