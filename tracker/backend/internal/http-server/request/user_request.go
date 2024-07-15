package request

import "tracker-app/backend/internal/models"

type AddUserRequest struct {
	PassportNumber string `json:"passportNumber" binding:"required"`
}

type GetUsersRequest struct {
	models.User
	Page    int `json:"page" form:"page" binding:"min=1"`
	PerPage int `json:"perPage" form:"perPage" binding:"min=1"`
}

type DeleteUserRequest struct {
	UserId models.UserId `binding:"required"`
}

type UpdateUserRequest struct {
	models.UserId `json:"userId"`
	Update        map[string]string `json:"update"`
}
