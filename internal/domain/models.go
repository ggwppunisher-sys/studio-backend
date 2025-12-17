package domain

import (
	"errors"
	"time"
)

type User struct {
	TgId             string // Telegram UserName
	TgChatId         string // Telegram chat id
	ListOfRecordings string // List of recordings
}

type Engineer struct {
	TgId             string    // Telegram UserName
	FreeTime         time.Time // Free time of Sound Engineer
	ListOfRecordings string    // List of recordings
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
	if u.TgId == "" {
		return errors.New("Incorrectly entered field")
	}
	if u.ListOfRecordings == "" {
		return errors.New("Incorrectly entered field")
	}
	return nil
}

func (e *Engineer) Validate() error {
	if e.TgId == "" {
		return errors.New("Incorrectly entered field")
	}
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
