package request

import "effective-mobile-golang/backend/internal/models"

type AddUserRequest struct {
	PassportNumber string `json:"passportNumber" binding:"required"`
}

type DeleteUserRequest struct {
	UserId int `json:"userId" binding:"required"`
}

type UpdateUserRequest struct {
	models.User
}
