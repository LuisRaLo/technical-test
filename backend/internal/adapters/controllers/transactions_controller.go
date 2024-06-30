package controllers

import (
	"technical-challenge/internal/core/domain/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ITransactionsController struct {
	Logger              *zap.SugaredLogger
	transactionsUseCase repositories.TransactionsUseCase
}

func NewTransactionsController(
	logger *zap.SugaredLogger,
	transactionsUseCase repositories.TransactionsUseCase,
) repositories.TransactionsController {
	logger.Info("TransactionsController created")
	return &ITransactionsController{
		Logger:              logger,
		transactionsUseCase: transactionsUseCase,
	}
}

// CreateTransaction implements repositories.TransactionsController.
// @Summary      Create a transaction
// @Description  Create a transaction
// @Tags        Transactions
// @Accept       json
// @Produce      json
// @Param        bond body repositories.CreateBondRequest true "Bond object"
// @Success      200  {object} repositories.CreateBondResponse200
// @Failure      400  {object} models.Response400WithResult
// @Failure      404  {object}  models.Response404WithResult
// @Failure      409  {object}  models.Response409WithResult
// @Failure      500  {object}  models.Response500WithResult
// @Router      /transactions [post]
// @Security    BearerAuth
func (i *ITransactionsController) CreateTransaction(transaction repositories.Transactions) error {
	panic("unimplemented")
}

// DeleteTransaction implements repositories.TransactionsController.
func (i *ITransactionsController) DeleteTransaction(id uuid.UUID) error {
	panic("unimplemented")
}

// GetTransactionByID implements repositories.TransactionsController.
func (i *ITransactionsController) GetTransactionByID(id uuid.UUID) (repositories.Transactions, error) {
	panic("unimplemented")
}

// GetTransactionsByBondID implements repositories.TransactionsController.
func (i *ITransactionsController) GetTransactionsByBondID(bondID uuid.UUID) ([]repositories.Transactions, error) {
	panic("unimplemented")
}

// GetTransactionsByUserID implements repositories.TransactionsController.
func (i *ITransactionsController) GetTransactionsByUserID(userID string) ([]repositories.Transactions, error) {
	panic("unimplemented")
}

// UpdateTransaction implements repositories.TransactionsController.
func (i *ITransactionsController) UpdateTransaction(transaction repositories.Transactions) error {
	panic("unimplemented")
}
