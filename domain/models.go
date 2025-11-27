package domain

import (
	"fmt"
	"time"
)

type User struct {
	Name        string
	Age         int
	DateOfBirth string
}

type Reservation struct {
	Name  string
	Place string
	Time  time.Time
}

type BookingScheduleEngineer struct {
	Name  string
	Place string
	Time  time.Time
}

type BookingScheduleUser struct {
	Name           string
	Place          string
	Time           time.Time
	WhoIsRecording string
}

func (u *User) Validate() error {
	if u.Name == "" {
		fmt.Println("Некорректно введенное имя")
	}
	if u.Age < 0 || u.Age > 100 {
		fmt.Println("Некорректно введенный возраст")
	}
	if u.DateOfBirth == "" {
		fmt.Println("Некорректно введенное поле рождения")
	}
	return nil
}

func (r *Reservation) Validate() error {
	if r.Name == "" {
		fmt.Println("Некорректно введенное имя")
	}
	if r.Place == "" {
		fmt.Println("Некорректно введенное место")
	}
	return nil
}

func (b *BookingScheduleEngineer) Validate() error {
	if b.Name == "" {
		fmt.Println("Некорректно введенное имя")
	}
	if b.Place == "" {
		fmt.Println("Некорректно введенное место")
	}
	return nil
}

func (bo *BookingScheduleUser) Validate() error {
	if bo.Name == "" {
		fmt.Println("Некорректно введенное имя")
	}
	if bo.Place == "" {
		fmt.Println("Некорректно введенное место")
	}
	if bo.WhoIsRecording == "" {
		fmt.Println("Некорректно введенное имя")
	}

	return nil
}
