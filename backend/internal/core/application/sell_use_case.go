package application

import (
	"net/http"
	"technical-challenge/internal/core/domain"
	"technical-challenge/internal/core/domain/models"

	"go.uber.org/zap"
)

type ISellUseCase struct {
	Logger *zap.SugaredLogger
}

func NewSellUseCase(
	logger *zap.SugaredLogger,
) domain.SellUseCase {
	logger.Info("SellUseCase created")
	return &ISellUseCase{
		Logger: logger,
	}
}

// Sell implements domain.SellUseCase.
func (i *ISellUseCase) Sell(request domain.SellRequest) models.DevResponse {
	return models.DevResponse{
		StatusCode: http.StatusOK,
		Response: models.Response{
			Message: "Operación realizada con éxito",
		},
	}
}
