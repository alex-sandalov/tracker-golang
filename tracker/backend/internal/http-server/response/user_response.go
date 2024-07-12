package response

import "tracker-app/backend/internal/models"

type AddUserResponse struct {
	models.User
}

type UpdateUserResponse struct {
	models.UserDB
}
