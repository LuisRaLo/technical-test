package repositories

import "github.com/google/uuid"

const (
	StatusBought = "BOUGHT"
	StatusSold   = "SOLD"
)

type (
	Transactions struct {
		ID        uuid.UUID `json:"id"`
		UserID    string    `json:"user_id"`
		BondID    uuid.UUID `json:"bond_id"`
		Status    string    `json:"status"`
		Quantity  int       `json:"quantity"`
		Price     float64   `json:"price"`
		CreatedAt int64     `json:"created_at"`
		UpdatedAt int64     `json:"updated_at"`
		DeletedAt int64     `json:"deleted_at,omitempty"`
	}

	TransactionsRepository interface {
		CreateTransaction(transaction Transactions) (int64, error)
		GetTransactionByID(id uuid.UUID) (Transactions, error)
		GetTransactionsByUserID(userID string) ([]Transactions, error)
		GetTransactionsByBondID(bondID uuid.UUID) ([]Transactions, error)
		GetTransactionsByStatusAndNotUserID(status string, userID string) ([]Transactions, error)
		GetTransactionsByUserIDAndStatus(userID string, status string) ([]Transactions, error)
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
