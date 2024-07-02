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

// GetAllBonds implements repositories.BondController.
// @Summary      Get all bonds
// @Description  Get all bonds
// @Tags        Bonds
// @Accept       json
// @Produce      json
// @Param        type path string true "Type of bond" Enums(SOLD, AVAILABLE, BOUGHT)
// @Success      200  {object} repositories.GetAllBondsResponse200
// @Failure      400  {object} models.Response400WithResult
// @Failure      404  {object}  models.Response404WithResult
// @Failure      409  {object}  models.Response409WithResult
// @Failure      500  {object}  models.Response500WithResult
// @Router      /bonds/{type} [get]
// @Security    BearerAuth
func (i *IBondsController) GetAllBonds() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context = r.Context()
		var userData utils.ResultFirebase = middlewares.GetUserDataFromContext(ctx)

		bondType := r.PathValue("type")

		var run models.DevResponse = i.bondUseCase.GetAllBonds(userData.TokenData.UID, bondType)

		utils.Response(w, run)
	}
}

// SellBond implements repositories.BondController.
// Sell a bound
// @Summary      Sell a bound
// @Description  Sell a bound
// @Tags         Sell
// @Accept       json
// @Produce      json
// @Param        sellRequest body domain.SellBondRequest true "Sell Request"
// @Success      200  {object}  domain.SellBondResponse200
// @Failure      400  {object} models.Response400WithResult
// @Failure      404  {object}  models.Response404WithResult
// @Failure      500  {object}  models.Response500WithResult
// @Router      /bonds/sell [post]
// @Security BearerAuth
func (i *IBondsController) SellBond() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context = r.Context()
		var payload repositories.SellBondRequest = repositories.SellBondRequest{}
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

		var run models.DevResponse = i.bondUseCase.SellBond(payload, userData.TokenData.UID)
		utils.Response(w, run)
	}
}
