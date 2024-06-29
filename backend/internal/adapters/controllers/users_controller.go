package controllers

import (
	"context"
	"net/http"
	"technical-challenge/internal/core/domain/constants"
	"technical-challenge/internal/core/domain/models"
	"technical-challenge/internal/core/domain/repositories"
	"technical-challenge/internal/middlewares"
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
// @Param        user body repositories.SignUpRequest true "User object that needs to be created"
// @Success      200  {object} repositories.UserResponse200
// @Failure      400  {object} models.Response400WithResult
// @Failure      404  {object}  models.Response404WithResult
// @Failure      500  {object}  models.Response500WithResult
// @Router      /users [post]
func (i *IUsersController) SingUp(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var signupRequest repositories.SignUpRequest = repositories.SignUpRequest{}

		errors := utils.ValidateBody(w, r, &signupRequest, i.Logger)
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

		emailValidation := utils.ValidateEmail(signupRequest.Email)
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

		passwordValidation := utils.ValidatePassword(signupRequest.Password, signupRequest.RetypedPassword)
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

		var run models.DevResponse = i.UserUseCase.CreateUser(ctx, signupRequest)

		utils.Response(w, run)
	}
}

// GetUserByID implements repositories.UserController.
func (i *IUsersController) GetUserByID(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("user_id")
		if userID == "" {
			utils.Response(w, models.DevResponse{
				StatusCode: http.StatusBadRequest,
				Response: models.Response400WithResult{
					Message: constants.REQUEST_INVALID,
					Details: []string{"User ID is required"},
				},
			})
			return
		}

		var run models.DevResponse = i.UserUseCase.GetUserByID(ctx, userID)

		utils.Response(w, run)
	}
}

// GetUserByJWT implements repositories.UserController.
func (i *IUsersController) GetUserByJWT() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context = r.Context()
		var userData utils.ResultFirebase = middlewares.GetUserDataFromContext(ctx)

		i.Logger.Info("User data: ", userData)

		var run models.DevResponse = i.UserUseCase.GetUserByID(ctx, userData.TokenData.UID)

		utils.Response(w, run)
	}
}
