package repositories

import (
	"net/http"
	"technical-challenge/internal/core/domain/models"

	"github.com/google/uuid"
)

type (
	Bond struct {
		ID        uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
		UserID    string    `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
		Name      string    `json:"name" example:"name"`
		Quantity  int       `json:"quantity" example:"1"`
		Price     float64   `json:"price" example:"1.00"`
		CreatedAt int64     `json:"created_at" example:"1618312800"`
		UpdatedAt int64     `json:"updated_at" example:"1618312800"`
		DeleteAt  int64     `json:"delete_at" example:"1618312800"`
	}

	BondModel struct {
		BondID            uuid.UUID `json:"bond_id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
		Seller            string    `json:"seller" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
		Name              string    `json:"name" example:"name"`
		TotalQuantity     int       `json:"total_quantity" example:"1"`
		Price             float64   `json:"price" example:"1.0000"`
		AvailableQuantity int       `json:"available_quantity" example:"1"`
	}

	BondSoldAndBought struct {
		ID            string  `json:"id"`
		Name          string  `json:"name"`
		Currency      string  `json:"currency"`
		NumerOfBonds  int     `json:"number_of_bonds"`
		TotalPrice    float64 `json:"total_price"`
		SellerOrBuyer string  `json:"seller_or_buyer"`
	}

	CreateBondRequest struct {
		Name     string  `json:"name" example:"name" validate:"required,min=3,max=40,alphanum"`
		Quantity int     `json:"quantity" example:"1" validate:"required,min=1,max=10000"`
		Price    float64 `json:"price" example:"1.0" validate:"required,min=0.0001,max=100000000.0000"`
	}

	CreateBondResponse struct {
		ID uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	}

	CreateBondResponse200 struct {
		*models.Response
		Result CreateBondResponse `json:"resultado"`
	}

	GetAllBondsResponse struct {
		Bonds []BondModel `json:"bonds"`
	}

	GetAllBondsSoldAndBoughtResponse struct {
		Bonds []BondSoldAndBought `json:"bonds"`
	}

	GetAllBondsResponse200 struct {
		*models.Response
		Result GetAllBondsResponse `json:"resultado"`
	}

	GetAllBondsSoldAndBoughtResponse200 struct {
		*models.Response
		Result GetAllBondsSoldAndBoughtResponse `json:"resultado"`
	}

	SellBondRequest struct {
		BondID   uuid.UUID `json:"bond_id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid" validate:"required,uuid"`
		Quantity int       `json:"quantity" example:"1" validate:"required,min=1,max=10000"`
	}

	SellBondResponse struct {
		ID uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	}

	SellBondResponse200 struct {
		*models.Response
		Result SellBondResponse `json:"resultado"`
	}

	BondRepository interface {
		CreateBond(bond Bond) int
		UpdateBond(bond Bond) int
		DeleteBond(id uuid.UUID) (int, error)
		GetBondByID(id uuid.UUID) (Bond, error)
		IsExistBond(name string, price float64, quantity int) bool
		GetAllBonds() ([]Bond, error)
		GetBondsSoldAndBought(userID string) ([]BondSoldAndBought, error)
		GetAllAvailableBonds(userID string) ([]BondModel, error)
		GetBondsByUserID(userID string) ([]Bond, error)
		GetBondByIDAndQuantity(id uuid.UUID, quantity int) (BondModel, error)
	}

	BondUseCase interface {
		CreateBond(bond CreateBondRequest, userID string) models.DevResponse
		GetAllBonds(userID string, bondType string) models.DevResponse
		SellBond(bond SellBondRequest, userID string) models.DevResponse
	}

	BondController interface {
		CreateBond() http.HandlerFunc
		GetAllBonds() http.HandlerFunc
		SellBond() http.HandlerFunc
	}
)
