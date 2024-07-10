package models

type User struct {
	UserId         int    `json:"userId" form:"user_id" binding:"required"`
	PassportNumber string `json:"passportNumber" form:"passport_number"`
	PasspoerSeries string `json:"passportSeries" form:"passport_series"`
	Surname        string `json:"surname" form:"surname"`
	Name           string `json:"name" form:"name"`
	Patronymic     string `json:"patronymic" form:"patronymic"`
	Address        string `json:"address" form:"address"`
}
