package service

import (
	"context"
	"info-golang/backend/internal/http-server/request"
	"info-golang/backend/internal/models"
	"info-golang/backend/internal/repository"
)

type UserService struct {
	repos repository.UserInterface
}

// NewUserService creates a new instance of UserService with the provided repository.
//
// repos: The repository.UserInterface implementation used to retrieve user information.
// Returns: A pointer to the newly created UserService instance.
func NewUserService(repos repository.UserInterface) *UserService {
	return &UserService{
		repos: repos,
	}
}

// GetUserInfo retrieves user information from the repository based on the provided user request.
//
// user: The request.GetInfoRequest object containing the user's passport series and number.
// Returns: The models.InfoUser object representing the retrieved user information
// and an error if the retrieval fails.
func (s *UserService) GetUserInfo(user request.GetInfoRequest) (models.InfoUser, error) {
	ctx := context.Background()

	return s.repos.GetUser(ctx, user)
}
