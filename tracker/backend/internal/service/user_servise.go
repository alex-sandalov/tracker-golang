package service

import (
	"context"
	"log/slog"
	"net/url"
	"tracker-app/backend/internal/http-server/request"
	"tracker-app/backend/internal/http-server/response"
	"tracker-app/backend/internal/lib"
	"tracker-app/backend/internal/models"
	"tracker-app/backend/internal/repository"

	"github.com/jmoiron/sqlx"
)

type UserService struct {
	log   *slog.Logger
	repos repository.UserInterface
	db    *sqlx.DB
}

func NewUserService(log *slog.Logger, repos repository.UserInterface, db *sqlx.DB) *UserService {
	return &UserService{
		log:   log,
		repos: repos,
		db:    db,
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
func (s *UserService) GetInfoUser(ctx context.Context, passportSeries, passportNumber string) (models.User, error) {
	var user models.User

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
// This function takes passport series and passport number as parameters and returns the ID of the newly added user and an error if the query fails.
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
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		s.log.Error("failed to begin transaction: %s", err)
		return models.UserId{}, err
	}

	id, err := s.repos.AddUser(ctx, tx, passportSeries, passportNumber)
	if err != nil {
		s.log.Error("failed to add user: %s", err)
		tx.Rollback()
		return models.UserId{}, err
	}

	err = tx.Commit()
	if err != nil {
		s.log.Error("failed to commit transaction: %s", err)
		return models.UserId{}, err
	}

	return id, nil
}

// DeleteUser deletes a user from the database.
//
// This function takes a DeleteUserRequest object containing the ID of the user to be deleted.
// It returns an error if the query fails.
//
// Parameters:
// - id: The DeleteUserRequest object containing the ID of the user to be deleted.
//
// Returns:
// - An error if the query fails.
func (s *UserService) DeleteUser(id request.DeleteUserRequest) error {
	ctx := context.Background()
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	err = s.repos.DeleteUser(ctx, tx, id.UserId)
	if err != nil {
		s.log.Error("failed to delete user: %s", err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		s.log.Error("failed to commit transaction: %s", err)
	}

	return err
}

// UpdateUser updates a user's information in the database.
//
// This function performs a transaction to update a user's information based on the provided UpdateUserRequest object. It first starts a transaction and then attempts to update the user. If the update is successful, it retrieves the updated user information to be included in the response. If any step fails, the transaction is rolled back, and an error is returned. On success, the transaction is committed, and the updated user information is returned.
//
// Parameters:
// - user: An UpdateUserRequest object containing the updated information for the user.
//
// Returns:
// - An UpdateUserResponse object containing the updated user information.
// - An error if the update fails at any step.
func (s *UserService) UpdateUser(user request.UpdateUserRequest) (response.UpdateUserResponse, error) {
	ctx := context.Background()
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return response.UpdateUserResponse{}, err
	}

	err = s.repos.UpdateUser(ctx, tx, user)
	if err != nil {
		s.log.Error("failed to update user: %s", err)
		tx.Rollback()
		return response.UpdateUserResponse{}, err
	}

	userInfo, err := s.repos.GetInfoUser(ctx, tx, user.UserId)
	if err != nil {
		s.log.Error("failed to get user info: %s", err)
		tx.Rollback()
		return response.UpdateUserResponse{}, err
	}

	err = tx.Commit()
	if err != nil {
		return response.UpdateUserResponse{}, err
	}

	return response.UpdateUserResponse{
		UserDB: userInfo,
	}, nil
}
