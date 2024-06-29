package repositories

import "github.com/google/uuid"

const (
	statusBought = "BOUGHT"
	statusSold   = "SOLD"
)

type (
	Transactions struct {
		ID        uuid.UUID `json:"id"`
		UserID    string    `json:"user_id"`
		BondID    uuid.UUID `json:"bond_id"`
		Status    string    `json:"status"`
		Quantity  int       `json:"quantity"`
		Price     float64   `json:"price"`
		CreatedAt int       `json:"created_at"`
		UpdatedAt int       `json:"updated_at"`
		DeletedAt int       `json:"deleted_at"`
	}

	TransactionsRepository interface {
		CreateTransaction(transaction Transactions) error
		GetTransactionByID(id uuid.UUID) (Transactions, error)
		GetTransactionsByUserID(userID string) ([]Transactions, error)
		GetTransactionsByBondID(bondID uuid.UUID) ([]Transactions, error)
		UpdateTransaction(transaction Transactions) error
		DeleteTransaction(id uuid.UUID) error
	}

	TransactionsUseCase interface {
		CreateTransaction(transaction Transactions) error
		GetTransactionByID(id uuid.UUID) (Transactions, error)
		GetTransactionsByUserID(userID string) ([]Transactions, error)
		GetTransactionsByBondID(bondID uuid.UUID) ([]Transactions, error)
		UpdateTransaction(transaction Transactions) error
		DeleteTransaction(id uuid.UUID) error
	}

	TransactionsController interface {
		CreateTransaction(transaction Transactions) error
		GetTransactionByID(id uuid.UUID) (Transactions, error)
		GetTransactionsByUserID(userID string) ([]Transactions, error)
		GetTransactionsByBondID(bondID uuid.UUID) ([]Transactions, error)
		UpdateTransaction(transaction Transactions) error
		DeleteTransaction(id uuid.UUID) error
	}
)
