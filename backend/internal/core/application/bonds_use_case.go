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

// GetAllBonds implements repositories.BondUseCase.
func (i *IBondsUseCase) GetAllBonds(userID string, bondType string) models.DevResponse {

	if bondType != "SOLD" && bondType != "AVAILABLE" && bondType != "BOUGHT" {
		return models.DevResponse{
			StatusCode: http.StatusBadRequest,
			Response: models.Response400WithResult{
				Message: "Invalid bond type",
				Details: []string{"Bond type must be SOLD, AVAILABLE or BOUGHT"},
			},
		}
	}

	if bondType == "AVAILABLE" {

		i.Logger.Info("=======================================================> AVAILABLE BONDS")

		bonds, err := i.bondRepository.GetAllAvailableBonds(userID)
		if err != nil {
			i.Logger.Error("Error getting all available bonds. ", err)
			return models.DevResponse{
				StatusCode: http.StatusInternalServerError,
				Response: models.Response500WithResult{
					Message: "ErrorGettingAllAvailableBonds",
				},
			}
		}

		var bondModels []repositories.BondModel
		for _, bond := range bonds {
			bondModels = append(bondModels, repositories.BondModel{
				BondID:            bond.BondID,
				Seller:            bond.Seller,
				Name:              bond.Name,
				TotalQuantity:     bond.TotalQuantity,
				Price:             bond.Price,
				AvailableQuantity: bond.AvailableQuantity,
			})
		}

		response200GetAllBonds := repositories.GetAllBondsResponse200{
			Response: &models.Response{
				Message: "Bonds retrieved",
			},
			Result: repositories.GetAllBondsResponse{
				Bonds: bondModels,
			},
		}

		return models.DevResponse{
			StatusCode: http.StatusOK,
			Response:   response200GetAllBonds,
		}
	}

	bonds, err := i.bondRepository.GetAllBondsBySOLDAndBOUGHT(bondType)
	if err != nil {
		i.Logger.Error("Error getting all bonds by type. ", err)
		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: "ErrorGettingAllBondsByType",
			},
		}
	}

	var bondModels []repositories.BondModel
	for _, bond := range bonds {
		bondModels = append(bondModels, repositories.BondModel{
			BondID:            bond.ID,
			Name:              bond.Name,
			TotalQuantity:     bond.Quantity,
			Price:             bond.Price,
			AvailableQuantity: 0,
		})
	}

	response200GetAllBonds := repositories.GetAllBondsResponse200{
		Response: &models.Response{
			Message: "Bonds retrieved",
		},
		Result: repositories.GetAllBondsResponse{
			Bonds: bondModels,
		},
	}

	return models.DevResponse{
		StatusCode: http.StatusOK,
		Response:   response200GetAllBonds,
	}
}
