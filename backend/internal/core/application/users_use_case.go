package application

import (
	"net/http"
	"technical-challenge/internal/core/domain/models"
	"technical-challenge/internal/core/domain/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IUsersUseCase struct {
	Logger          *zap.SugaredLogger
	usersRepository repositories.UsersRepository
}

func NewUsersUseCase(
	logger *zap.SugaredLogger,
	usersRepository repositories.UsersRepository,
) repositories.UserUseCase {
	logger.Info("UsersUseCase created")
	return &IUsersUseCase{
		Logger:          logger,
		usersRepository: usersRepository,
	}
}

// CreateUser implements repositories.UserUseCase.
func (i *IUsersUseCase) CreateUser(user *repositories.User) models.DevResponse {

	return models.DevResponse{
		StatusCode: http.StatusCreated,
		Response:   "OK",
	}
}

// GetUserByEmail implements repositories.UserUseCase.
func (i *IUsersUseCase) GetUserByEmail(email string) models.DevResponse {
	panic("unimplemented")
}

// GetUserByID implements repositories.UserUseCase.
func (i *IUsersUseCase) GetUserByID(id uuid.UUID) models.DevResponse {
	panic("unimplemented")
}
