package service

import (
	"context"
	"log/slog"
	"net/url"
	"tracker-app/backend/internal/http-server/request"
	"tracker-app/backend/internal/lib"
	"tracker-app/backend/internal/models"
	"tracker-app/backend/internal/repository"
)

type UserService struct {
	log   *slog.Logger
	repos repository.UserInterface
}

func NewUserService(log *slog.Logger, repos repository.UserInterface) *UserService {
	return &UserService{
		log:   log,
		repos: repos,
	}
}

// GetInfoUser retrieves user information from the external API based on the user's passport series and number.
//
// Parameters:
// - ctx: The context for the HTTP request.
// - idUserTracker: The internal user ID.
// - passportSeries: The passport series.
// - passportNumber: The passport number.
//
// Returns:
// - The user information.
// - An error if the request fails.
func (s *UserService) GetInfoUser(ctx context.Context, idUserTracker models.UserId, passportSeries, passportNumber string) (models.User, error) {
	var user models.User

	user.UserId = idUserTracker

	user.PassportNumber = passportNumber
	user.PasspoerSeries = passportSeries

	params := url.Values{}
	params.Add("passportSerie", passportSeries)
	params.Add("passportNumber", passportNumber)

	url := lib.RequestFormat("http://info-golang:8081/api/internal/info", params)

	infoUser, err := lib.GET(ctx, url)
	if err != nil {
		s.log.Error("failed to get user info: %s", err)
		return user, err
	}

	user.Surname = infoUser.Surname
	user.Name = infoUser.Name
	user.Patronymic = infoUser.Patronymic
	user.Address = infoUser.Address

	return user, nil
}

// AddUser adds a new user to the database.
//
// Parameters:
// - passportSeries: The passport series.
// - passportNumber: The passport number.
//
// Returns:
// - The ID of the newly added user.
// - An error if the query fails.
func (s *UserService) AddUser(passportSeries, passportNumber string) (models.UserId, error) {
	ctx := context.Background()

	return s.repos.AddUser(ctx, passportSeries, passportNumber)
}

// DeleteUser deletes a user from the database.
//
// Parameters:
// - id: The ID of the user to be deleted.
//
// Returns:
// - An error if the query fails.
func (s *UserService) DeleteUser(id request.DeleteUserRequest) error {
	ctx := context.Background()

	return s.repos.DeleteUser(ctx, id.UserId)
}

func (s *UserService) UpdateUser(user request.UpdateUserRequest) error {
	return nil
}
