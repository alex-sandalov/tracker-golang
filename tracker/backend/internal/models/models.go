package models

type User struct {
	UserId
	PassportNumber string `json:"passportNumber" form:"passportNumber"`
	PasspoerSeries string `json:"passportSeries" form:"passportSeries"`
	Surname        string `json:"surname" form:"surname"`
	Name           string `json:"name" form:"name"`
	Patronymic     string `json:"patronymic" form:"patronymic"`
	Address        string `json:"address" form:"address"`
}

type UserDB struct {
	UserId         int64 `db:"user_id"`
	PassportNumber string `db:"passport_number"`
	PasspoerSerie  string `db:"passport_serie"`
}

type UserId struct {
	UserId int64 `json:"userId" form:"userId" db:"user_id"`
}
