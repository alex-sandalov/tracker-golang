package response

import "info-golang/backend/internal/models"

type GetUserInfoResponse struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

// NewGetUserInfoResponse creates a new GetUserInfoResponse object from a models.InfoUser object.
//
// user: The models.InfoUser object to convert.
// Returns: A new GetUserInfoResponse object with the user's information.
func NewGetUserInfoResponse(user models.InfoUser) GetUserInfoResponse {
	response := GetUserInfoResponse{
		Surname: user.Surname,
		Name:    user.Name,
		Address: user.Address,
	}

	if user.Patronymic == "" {
		response.Patronymic = "не указано"
	} else {
		response.Patronymic = user.Patronymic
	}

	return response
}
