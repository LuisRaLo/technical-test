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

type IBondsController struct {
	Logger      *zap.SugaredLogger
	bondUseCase repositories.BondUseCase
}

func NewBondsController(
	logger *zap.SugaredLogger,
	bondUseCase repositories.BondUseCase,
) repositories.BondController {
	logger.Info("BondsController created")
	return &IBondsController{
		Logger:      logger,
		bondUseCase: bondUseCase,
	}
}

// CreateBond implements repositories.BondController.
// @Summary      Create a bond
// @Description  Create a bond
// @Tags        Bonds
// @Accept       json
// @Produce      json
// @Param        bond body repositories.CreateBondRequest true "Bond object"
// @Success      200  {object} repositories.CreateBondResponse200
// @Failure      400  {object} models.Response400WithResult
// @Failure      404  {object}  models.Response404WithResult
// @Failure      409  {object}  models.Response409WithResult
// @Failure      500  {object}  models.Response500WithResult
// @Router      /bonds [post]
// @Security    BearerAuth
func (i *IBondsController) CreateBond() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context = r.Context()
		var payload repositories.CreateBondRequest = repositories.CreateBondRequest{}
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

		var run models.DevResponse = i.bondUseCase.CreateBond(payload, userData.TokenData.UID)
		utils.Response(w, run)
	}
}
