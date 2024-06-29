package repositories

import (
	"net/http"
	"technical-challenge/internal/core/domain/models"

	"github.com/google/uuid"
)

type (
	User struct {
		ID        uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
		Email     string    `json:"email" example:"email@email.com"`
		CreatedAt int64     `json:"created_at" example:"1618312800"`
		UpdatedAt int64     `json:"updated_at" example:"1618312800"`
		DeleteAt  int64     `json:"delete_at" example:"1618312800"`
	}

	UsersRepository interface {
		CreateUser(user *User) error
		GetUserByID(id uuid.UUID) (*User, error)
		GetUserByEmail(email string) (*User, error)
	}

	UserUseCase interface {
		CreateUser(user *User) models.DevResponse
		GetUserByID(id uuid.UUID) models.DevResponse
		GetUserByEmail(email string) models.DevResponse
	}

	UserController interface {
		CreateUser() http.HandlerFunc
		GetUserByID() http.HandlerFunc
		GetUserByEmail() http.HandlerFunc
	}
)
