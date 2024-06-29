package controllers

import (
	"context"
	"net/http"
	"technical-challenge/internal/core/domain"
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

// @Failure      400  {object} models.Response400WithResult
// @Failure      404  {object}  models.Response404WithResult
// @Failure      500  {object}  models.Response500WithResult
// @Router      /users [post]

func (i *IUsersController) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context = r.Context()
		var payload domain.SellRequest = domain.SellRequest{}
		var userData utils.ResultFirebase = middlewares.GetUserDataFromContext(ctx)

		errors := utils.ValidateBody(w, r, &payload, i.Logger)
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

		//var run models.DevResponse = i.SellUseCase.Sell(userData.TokenData.UID, payload)
		utils.Response(w, models.DevResponse{
			StatusCode: http.StatusOK,
			Response: models.ResponseWithResult{
				Result: userData,
			},
		})
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
