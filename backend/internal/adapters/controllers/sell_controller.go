package controllers

import (
	"context"
	"net/http"
	"technical-challenge/internal/core/domain"
	"technical-challenge/internal/core/domain/constants"
	"technical-challenge/internal/core/domain/models"
	"technical-challenge/internal/middlewares"
	"technical-challenge/internal/utils"

	"go.uber.org/zap"
)

type ISellController struct {
	Logger      *zap.SugaredLogger
	SellUseCase domain.SellUseCase
}

func NewSellController(
	logger *zap.SugaredLogger,
	sellUseCase domain.SellUseCase,
) domain.SellController {
	logger.Info("SellController created")
	return &ISellController{
		Logger:      logger,
		SellUseCase: sellUseCase,
	}
}

// Sell godoc
// @Summary      Sell a bound
// @Description  Sell a bound
// @Tags         Sell
// @Accept       json
// @Produce      json
// @Param        sellRequest body domain.SellRequest true "Sell Request"
// @Success      200  {object}  domain.SellResponse
// @Failure      400  {object} models.Response400WithResult
// @Failure      404  {object}  models.Response404WithResult
// @Failure      500  {object}  models.Response500WithResult
// @Router      /sell [post]
// @Security BearerAuth
func (i *ISellController) Sell() http.HandlerFunc {
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

		var run models.DevResponse = i.SellUseCase.Sell(userData.TokenData.UID, payload)
		utils.Response(w, run)
	}

}
