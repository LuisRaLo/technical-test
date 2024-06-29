package application

import (
	"context"
	"net/http"
	"technical-challenge/internal/core/domain/constants"
	"technical-challenge/internal/core/domain/models"
	"technical-challenge/internal/core/domain/repositories"
	"technical-challenge/internal/utils/helpers"

	"firebase.google.com/go/auth"

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
func (i *IUsersUseCase) CreateUser(ctx context.Context, user repositories.SignUpRequest) models.DevResponse {
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

	response200CreateUser := repositories.UserResponse200{
		Response: &models.Response{
			Message: "User created",
		},
		Result: repositories.UserModel{
			ID:        userModel.ID,
			Email:     userModel.Email,
			Name:      userModel.Name,
			CreatedAt: userModel.CreatedAt,
			UpdatedAt: userModel.UpdatedAt,
		},
	}

	return models.DevResponse{
		StatusCode: http.StatusOK,
		Response:   response200CreateUser,
	}
}

// GetUserByID implements repositories.UserUseCase.
func (i *IUsersUseCase) GetUserByID(ctx context.Context, id string) models.DevResponse {
	i.Logger.Info("Getting user by ID: ", id)

	user, err := i.usersRepository.GetUserByID(id)
	if err != nil {
		i.Logger.Error("Error getting user by ID")
		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: "Error getting user by ID",
			},
		}
	}

	if user.ID == "" {
		return models.DevResponse{
			StatusCode: http.StatusNotFound,
			Response: models.Response404WithResult{
				Message: constants.REQUEST_USER_NOT_FOUND,
				Details: []string{"User not found"},
			},
		}
	}

	response200GetUserByID := repositories.UserResponse200{
		Response: &models.Response{
			Message: "User found",
		},
		Result: repositories.UserModel{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	return models.DevResponse{
		StatusCode: http.StatusOK,
		Response:   response200GetUserByID,
	}
}
