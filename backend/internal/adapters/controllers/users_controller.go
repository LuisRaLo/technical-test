package controllers

import (
	"context"
	"net/http"
	"technical-challenge/internal/core/domain/constants"
	"technical-challenge/internal/core/domain/models"
	"technical-challenge/internal/core/domain/repositories"
	"technical-challenge/internal/utils"

	"go.uber.org/zap"
)

type IUsersController struct {
	Logger      *zap.SugaredLogger
	UserUseCase repositories.UserUseCase
}

func NewUsersController(
	logger *zap.SugaredLogger,
	userUseCase repositories.UserUseCase,
) repositories.UserController {
	logger.Info("UsersController created")
	return &IUsersController{
		Logger:      logger,
		UserUseCase: userUseCase,
	}
}

// CreateUser implements repositories.UserController.
// @Summary      Create a user
// @Description  Create a user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user body repositories.CreateUserRequest true "User object that needs to be created"
// @Success      200  {object} repositories.Response200CreateUser
// @Failure      400  {object} models.Response400WithResult
// @Failure      404  {object}  models.Response404WithResult
// @Failure      500  {object}  models.Response500WithResult
// @Router      /users [post]
func (i *IUsersController) CreateUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createUserRequest repositories.CreateUserRequest = repositories.CreateUserRequest{}

		errors := utils.ValidateBody(w, r, &createUserRequest, i.Logger)
		if len(errors) > 0 {
			utils.Response(w, models.DevResponse{
				StatusCode: http.StatusBadRequest,
				Response: models.Response400WithResult{
					Message: constants.REQUEST_INVALID,
					Details: errors,
				},
			})
			return
		}

		emailValidation := utils.ValidateEmail(createUserRequest.Email)
		if len(emailValidation) > 0 {
			utils.Response(w, models.DevResponse{
				StatusCode: http.StatusBadRequest,
				Response: models.Response400WithResult{
					Message: constants.REQUEST_INVALID,
					Details: []string{emailValidation},
				},
			})
			return
		}

		passwordValidation := utils.ValidatePassword(createUserRequest.Password, createUserRequest.RetypedPassword)
		if len(passwordValidation) > 0 {
			utils.Response(w, models.DevResponse{
				StatusCode: http.StatusBadRequest,
				Response: models.Response400WithResult{
					Message: constants.REQUEST_INVALID,
					Details: []string{passwordValidation},
				},
			})
			return
		}

		var run models.DevResponse = i.UserUseCase.CreateUser(ctx, createUserRequest)

		utils.Response(w, run)
	}
}

// GetUserByEmail implements repositories.UserController.
func (i *IUsersController) GetUserByEmail() http.HandlerFunc {
	panic("unimplemented")
}

// GetUserByID implements repositories.UserController.
func (i *IUsersController) GetUserByID() http.HandlerFunc {
	panic("unimplemented")
}
