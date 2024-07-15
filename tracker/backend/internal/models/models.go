package models

import "time"

type User struct {
	UserId
	PassportNumber string `json:"passportNumber" form:"passportNumber" db:"passport_number"`
	PassportSeries string `json:"passportSeries" form:"passportSerie" db:"passport_serie"`
	Surname        string `json:"surname" form:"surname"`
	Name           string `json:"name" form:"name"`
	Patronymic     string `json:"patronymic" form:"patronymic"`
	Address        string `json:"address" form:"address"`
}

type UserDB struct {
	UserId         int64  `db:"user_id"`
	PassportNumber string `db:"passport_number"`
	PasspoerSerie  string `db:"passport_serie"`
}

type UserId struct {
	UserId int64 `json:"userId" form:"userId" db:"user_id"`
}

type Task struct {
	UserId
	TaskId      int64     `json:"taskId" form:"taskId" db:"task_id"`
	Description string    `json:"description" form:"description" db:"description"`
	TimeStart   time.Time `json:"timeStart" form:"timeStart" db:"start_time"`
	TimeStop    time.Time `json:"timeStop" form:"timeStop" db:"end_time"`
	Active      bool      `json:"active" form:"active" db:"active"`
	Duration    []uint8   `db:"duration"`
}
