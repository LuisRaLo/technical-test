package repositories

import (
	"context"
	"net/http"
	"technical-challenge/internal/core/domain/models"

	"github.com/google/uuid"
)

type (
	User struct {
		ID        string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
		Email     string `json:"email" example:"email@email.com"`
		Name      string `json:"name" example:"name"`
		CreatedAt int64  `json:"created_at" example:"1618312800"`
		UpdatedAt int64  `json:"updated_at" example:"1618312800"`
		DeleteAt  int64  `json:"delete_at,omitempty" example:"1618312800"`
	}

	CreateUserRequest struct {
		Name            string `json:"name" example:"name" validate:"required,max=255"`
		Email           string `json:"email" example:"email@email.com" validate:"required,email,max=255" format:"email"`
		Password        string `json:"password" example:"password " validate:"required,min=6"`
		RetypedPassword string `json:"retyped_password" example:"password" validate:"required,min=6"`
	}

	Response200CreateUser struct {
		*models.Response
		Result User `json:"result"`
	}

	UsersRepository interface {
		CreateUser(user *User) int
		GetUserByID(id uuid.UUID) (*User, error)
		GetUserByEmail(email string) (*User, error)
		IsEmailAlreadyInUse(email string) bool
	}

	UserUseCase interface {
		CreateUser(ctx context.Context, user CreateUserRequest) models.DevResponse
		GetUserByID(id uuid.UUID) models.DevResponse
		GetUserByEmail(email string) models.DevResponse
	}

	UserController interface {
		CreateUser(ctx context.Context) http.HandlerFunc
		GetUserByID() http.HandlerFunc
		GetUserByEmail() http.HandlerFunc
	}
)
