package repositories

import "github.com/google/uuid"

type (
	Transaction struct {
		ID         int64     `json:"id" example:"1"`
		OprationID int64     `json:"opration_id" example:"1"`
		UserID     uuid.UUID `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
		CreatedAt  int64     `json:"created_at" example:"1618312800"`
		UpdatedAt  int64     `json:"updated_at" example:"1618312800"`
		DeleteAt   int64     `json:"delete_at" example:"1618312800"`
	}
)
