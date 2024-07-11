package service

import (
	"info-golang/backend/internal/http-server/request"
	"info-golang/backend/internal/models"
	"info-golang/backend/internal/repository"
)

type UserInterface interface {
	GetUserInfo(user request.GetInfoRequest) (models.InfoUser, error)
}

type Service struct {
	UserInterface
}

// NewService creates a new Service instance with the provided repository.
//
// Parameters:
// - repos: The repository.Repository instance containing the repository implementations.
//
// Returns:
// - *Service: A pointer to the newly created Service instance.
func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserInterface: NewUserService(repos.UserInterface),
	}
}
