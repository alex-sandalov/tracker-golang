package models

type InfoUser struct {
	UserId         int    `db:"id"`
	PassportNumber int    `db:"passport_number"`
	PassportSerie  int    `db:"passport_serie"`
	Surname        string `json:"surname" db:"surname"`
	Name           string `json:"name" db:"name"`
	Patronymic     string `json:"patronymic" db:"patronymic"`
	Address        string `json:"address" db:"address"`
}
