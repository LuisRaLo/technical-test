package domain

import (
	"net/http"
	"technical-challenge/internal/core/domain/models"
)

type (
	SellUseCase interface {
		Sell(useID string, request SellRequest) models.DevResponse
	}

	SellController interface {
		Sell() http.HandlerFunc
	}

	SellRequest struct {
		Name     string  `json:"name" example:"name" validate:"required,min=3,max=40,alphanum"`
		Quantity int     `json:"quantity" example:"1" validate:"required,min=1,max=10000"`
		Price    float64 `json:"price" example:"1.0" validate:"required,min=0.0001,max=100000000.0000"`
	}

	SellResponseResult struct {
		ID string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	}
	SellResponse struct {
		*models.Response
		Result SellResponseResult `json:"resultado"`
	}

	SellError struct {
		Code    int    `json:"code" example:"400"`
		Message string `json:"message" example:"error message"`
	}
)
