package repositories

import (
	"context"
	"database/sql"
	"net/http"
	"technical-challenge/internal/core/domain/models"
)

type (
	User struct {
		ID        string        `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
		Email     string        `json:"email" example:"email@email.com"`
		Name      string        `json:"name" example:"name"`
		CreatedAt int64         `json:"created_at" example:"1618312800"`
		UpdatedAt int64         `json:"updated_at" example:"1618312800"`
		DeleteAt  sql.NullInt64 `json:"delete_at,omitempty" example:"1618312800"`
	}

	UserModel struct {
		ID        string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
		Email     string `json:"email" example:"email@email.com"`
		Name      string `json:"name" example:"name"`
		CreatedAt int64  `json:"created_at" example:"1618312800"`
		UpdatedAt int64  `json:"updated_at" example:"1618312800"`
	}

	SignUpRequest struct {
		Name            string `json:"name" example:"name" validate:"required,max=255"`
		Email           string `json:"email" example:"email@email.com" validate:"required,email,max=255" format:"email"`
		Password        string `json:"password" example:"password " validate:"required,min=6"`
		RetypedPassword string `json:"retyped_password" example:"password" validate:"required,min=6"`
	}

	UserResponse200 struct {
		*models.Response
		Result UserModel `json:"result"`
	}

	SingInRequest struct {
		Email    string `json:"email" example:"email@email.com" validate:"required,email,max=255" format:"email"`
		Password string `json:"password" example:"password " validate:"required,min=6"`
	}

	UsersRepository interface {
		CreateUser(user *User) int
		GetUserByID(id string) (User, error)
		IsEmailAlreadyInUse(email string) bool
	}

	UserUseCase interface {
		CreateUser(ctx context.Context, user SignUpRequest) models.DevResponse
		GetUserByID(ctx context.Context, id string) models.DevResponse
	}

	UserController interface {
		SingUp(ctx context.Context) http.HandlerFunc
		GetUserByJWT() http.HandlerFunc
		GetUserByID(ctx context.Context) http.HandlerFunc
	}
)
