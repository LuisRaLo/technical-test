package repositories

type (
	OperationsCatalog struct {
		ID          int64  `json:"id" example:"1"`
		Name        string `json:"name" example:"name"`
		Description string `json:"description" example:"description"`
		CreatedAt   int64  `json:"created_at" example:"1618312800"`
		UpdatedAt   int64  `json:"updated_at" example:"1618312800"`
		DeleteAt    int64  `json:"delete_at" example:"1618312800"`
	}
)
