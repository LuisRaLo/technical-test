package repositories

import (
	"technical-challenge/internal/core/domain/models"

	"github.com/google/uuid"
)

type (
	Bound struct {
		ID        uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
		UserID    uuid.UUID `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
		Name      string    `json:"name" example:"name"`
		Quantity  int       `json:"quantity" example:"1"`
		Price     float64   `json:"price" example:"1.00"`
		CreatedAt int64     `json:"created_at" example:"1618312800"`
		UpdatedAt int64     `json:"updated_at" example:"1618312800"`
		DeleteAt  int64     `json:"delete_at" example:"1618312800"`
	}

	CreateBoundRequest struct {
		Name     string  `json:"name" example:"name" validate:"required,min=3,max=40,alphanum"`
		Quantity int     `json:"quantity" example:"1" validate:"required,min=1,max=10000"`
		Price    float64 `json:"price" example:"1.0" validate:"required,min=0.0001,max=100000000.0000"`
	}

	BoundRepository interface {
		CreateBound(bound Bound) error
		UpdateBound(bound Bound) error
		DeleteBound(bound Bound) error
		GetBoundByID(id uuid.UUID) (Bound, error)
		GetAllBounds() ([]Bound, error)
	}

	BoundUseCase interface {
		CreateBound(bound CreateBoundRequest) models.DevResponse
	}

	BoundController interface {
		CreateBound() models.DevResponse
	}
)
