package application

import (
	"technical-challenge/internal/core/domain/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ITransactionsUseCase struct {
	Logger                 *zap.SugaredLogger
	transactionsRepository repositories.TransactionsRepository
}

func NewTransactionsUseCase(
	logger *zap.SugaredLogger,
	transactionsRepository repositories.TransactionsRepository,
) repositories.TransactionsUseCase {
	logger.Info("TransactionsUseCase created")
	return &ITransactionsUseCase{
		Logger:                 logger,
		transactionsRepository: transactionsRepository,
	}
}

// CreateTransaction implements repositories.TransactionsUseCase.
func (i *ITransactionsUseCase) CreateTransaction(transaction repositories.Transactions) error {
	panic("unimplemented")
}

// DeleteTransaction implements repositories.TransactionsUseCase.
func (i *ITransactionsUseCase) DeleteTransaction(id uuid.UUID) error {
	panic("unimplemented")
}

// GetTransactionByID implements repositories.TransactionsUseCase.
func (i *ITransactionsUseCase) GetTransactionByID(id uuid.UUID) (repositories.Transactions, error) {
	panic("unimplemented")
}

// GetTransactionsByBondID implements repositories.TransactionsUseCase.
func (i *ITransactionsUseCase) GetTransactionsByBondID(bondID uuid.UUID) ([]repositories.Transactions, error) {
	panic("unimplemented")
}

// GetTransactionsByUserID implements repositories.TransactionsUseCase.
func (i *ITransactionsUseCase) GetTransactionsByUserID(userID string) ([]repositories.Transactions, error) {
	panic("unimplemented")
}

// UpdateTransaction implements repositories.TransactionsUseCase.
func (i *ITransactionsUseCase) UpdateTransaction(transaction repositories.Transactions) error {
	panic("unimplemented")
}
