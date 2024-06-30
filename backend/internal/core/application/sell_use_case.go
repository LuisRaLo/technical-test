package application

import (
	"net/http"
	"technical-challenge/internal/core/domain"
	"technical-challenge/internal/core/domain/models"
	"technical-challenge/internal/core/domain/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ISellUseCase struct {
	Logger         *zap.SugaredLogger
	bondRepository repositories.BondRepository
}

func NewSellUseCase(
	logger *zap.SugaredLogger,
	bondRepository repositories.BondRepository,
) domain.SellUseCase {
	logger.Info("SellUseCase created")
	return &ISellUseCase{
		Logger:         logger,
		bondRepository: bondRepository,
	}
}

// Sell implements domain.SellUseCase.
func (i *ISellUseCase) Sell(useID string, request domain.SellRequest) models.DevResponse {
	var sellResponse domain.SellResponse = domain.SellResponse{}

	//generate a random id
	bondId := uuid.New().String()

	//TODO: comprobar que el usuario este dentro de las txs permitidas
	//TODO: comprobar que el usuario tenga el saldo suficiente
	//TODO: comprobar que el bono exista, tenga la cantidad suficiente y no este vendido
	//TODO: realizar la venta
	//TODO: actualizar el saldo del usuario
	//TODO: actualizar la cantidad del bono
	//TODO: actualizar el estado del bono
	//TODO: guardar la transaccion
	//TODO: guardar la venta

	//creamos respuesta exitosa
	sellResponse.Result.ID = bondId

	return models.DevResponse{
		StatusCode: http.StatusOK,
		Response:   sellResponse,
	}
}
