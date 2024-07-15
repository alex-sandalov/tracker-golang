package response

import "tracker-app/backend/internal/models"

type AddUserResponse struct {
	models.User
}

type GetUserResponse struct {
	models.UserId
	PassportSerie  string `json:"passportSerie"`
	PassportNumber string `json:"passportNumber"`
}

type GetUsersResponse struct {
	CountUsersAll  int               `json:"counUsersAll"`
	CountUsersPage int               `json:"countUsersPage"`
	Users          []GetUserResponse `json:"users"`
}

type UpdateUserResponse struct {
	models.UserDB
}
