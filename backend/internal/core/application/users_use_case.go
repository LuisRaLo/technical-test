package application

import (
	"context"
	"net/http"
	"technical-challenge/internal/core/domain/constants"
	"technical-challenge/internal/core/domain/models"
	"technical-challenge/internal/core/domain/repositories"
	"technical-challenge/internal/utils/helpers"

	"firebase.google.com/go/auth"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IUsersUseCase struct {
	Logger          *zap.SugaredLogger
	usersRepository repositories.UsersRepository
	fireaseConn     *auth.Client
}

func NewUsersUseCase(
	logger *zap.SugaredLogger,
	usersRepository repositories.UsersRepository,
	fireaseConn *auth.Client,
) repositories.UserUseCase {
	logger.Info("UsersUseCase created")
	return &IUsersUseCase{
		Logger:          logger,
		usersRepository: usersRepository,
		fireaseConn:     fireaseConn,
	}
}

// CreateUser implements repositories.UserUseCase.
func (i *IUsersUseCase) CreateUser(ctx context.Context, user repositories.CreateUserRequest) models.DevResponse {
	i.Logger.Info("Creating user with payload: ", user)

	isEmailInUse := i.usersRepository.IsEmailAlreadyInUse(user.Email)
	if isEmailInUse {
		return models.DevResponse{
			StatusCode: http.StatusBadRequest,
			Response: models.Response400WithResult{
				Message: constants.REQUEST_INVALID,
				Details: []string{"Email already in use"},
			},
		}
	}

	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(false).
		Password(user.Password).
		DisplayName(user.Name)

	// Create user in firebase
	u, err := i.fireaseConn.CreateUser(ctx, params)
	if err != nil {
		i.Logger.Error("Error creating user in firebase")
		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: "Error creating user in firebase",
			},
		}
	}

	timestamp := helpers.GetTimeNow()
	var userModel repositories.User = repositories.User{
		ID:        u.UID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}

	// Create user in database
	idInserted := i.usersRepository.CreateUser(&userModel)
	if idInserted == 0 {
		i.Logger.Error("Error creating user in database")

		// Delete user from firebase
		errDelete := i.fireaseConn.DeleteUser(ctx, u.UID)
		if errDelete != nil {
			i.Logger.Error("Error deleting user from firebase")
		}

		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: "Error creating user in database",
			},
		}
	}

	response200CreateUser := repositories.Response200CreateUser{
		Response: &models.Response{
			Message: "User created",
		},
		Result: userModel,
	}

	return models.DevResponse{
		StatusCode: http.StatusOK,
		Response:   response200CreateUser,
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
