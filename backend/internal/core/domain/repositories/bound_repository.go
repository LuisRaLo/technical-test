package repositories

import "github.com/google/uuid"

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
)
