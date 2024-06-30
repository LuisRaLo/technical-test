package application

import (
	"net/http"
	"technical-challenge/internal/core/domain/models"
	"technical-challenge/internal/core/domain/repositories"
	"technical-challenge/internal/utils/helpers"

	"firebase.google.com/go/auth"
	"github.com/google/uuid"

	"go.uber.org/zap"
)

type IBondsUseCase struct {
	Logger         *zap.SugaredLogger
	bondRepository repositories.BondRepository
	fireaseConn    *auth.Client
}

func NewBondsUseCase(
	logger *zap.SugaredLogger,
	bondRepository repositories.BondRepository,
	fireaseConn *auth.Client,
) repositories.BondUseCase {
	logger.Info("BondsUseCase created")
	return &IBondsUseCase{
		Logger:         logger,
		bondRepository: bondRepository,
		fireaseConn:    fireaseConn,
	}
}

// CreateBond implements repositories.BondUseCase.
func (i *IBondsUseCase) CreateBond(payload repositories.CreateBondRequest, userID string) models.DevResponse {

	isExists := i.bondRepository.IsExistBond(payload.Name, payload.Price, payload.Quantity)
	if isExists {
		i.Logger.Error("Bond already exists")
		return models.DevResponse{
			StatusCode: http.StatusConflict,
			Response: models.Response409WithResult{
				Message: "Bond already exists",
			},
		}
	}
	timestamp := helpers.GetTimeNow()
	bondID := uuid.New()

	r := i.bondRepository.CreateBond(repositories.Bond{
		ID:        bondID,
		UserID:    userID,
		Name:      payload.Name,
		Quantity:  payload.Quantity,
		Price:     payload.Price,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
		DeleteAt:  0,
	})

	if r == 0 {
		i.Logger.Error("Error creating bond. ")

		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: "ErrorCreatingBond",
			},
		}
	}

	bond, err := i.bondRepository.GetBondByID(bondID)
	if err != nil {
		i.Logger.Error("Error getting bond by ID. ", err)

		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: "ErrorGettingBond",
			},
		}
	}

	response200CreateBond := repositories.CreateBondResponse200{
		Response: &models.Response{
			Message: "Bond created",
		},
		Result: repositories.CreateBondResponse{
			ID: bond.ID,
		},
	}

	return models.DevResponse{
		StatusCode: http.StatusOK,
		Response:   response200CreateBond,
	}

}
