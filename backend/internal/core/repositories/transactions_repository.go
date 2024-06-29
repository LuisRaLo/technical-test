package repositories

import (
	"database/sql"
	"io"
	"os"
	"technical-challenge/internal/core/domain/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ITransactionRepository struct {
	Logger     *zap.SugaredLogger
	datasoruce *sql.DB
}

func NewTransactionRepository(
	logger *zap.SugaredLogger,
	datasoruce *sql.DB,
) repositories.TransactionsRepository {
	logger.Info("BoundRepository created")

	ITransactionRepository := &ITransactionRepository{
		Logger:     logger,
		datasoruce: datasoruce,
	}

	if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "local" {
		logger.Info("Syncing database")
		ITransactionRepository.SyncDatabase()
	}

	return ITransactionRepository
}

// Función para sincronizar la base de datos ejecutando un archivo SQL.
func (i *ITransactionRepository) SyncDatabase() {
	// Leer el contenido del archivo SQL
	var filepath string = "test/migrations/transacctions.sql"

	file, err := os.Open(filepath)
	if err != nil {
		i.Logger.Error("Error opening file")
	}

	defer file.Close()

	// Read the content of the SQL file
	content, err := io.ReadAll(file)
	if err != nil {
		i.Logger.Error("Error reading file")
	}

	//imprimir el contenido del archivo
	i.Logger.Info(string(content))

	// Ejecutar el contenido del archivo SQL
	_, err = i.datasoruce.Exec(string(content))
	if err != nil {
		i.Logger.Error("Error executing SQL file. ", err)
	}

	i.Logger.Info("Database synced")
}

// Exec
func (i *ITransactionRepository) Exec(query string) {
	_, err := i.datasoruce.Exec(query)
	if err != nil {
		i.Logger.Error("Error executing query. ", err)
	}
}

// CreateTransaction implements repositories.TransactionsRepository.
func (i *ITransactionRepository) CreateTransaction(transaction repositories.Transactions) error {
	panic("unimplemented")
}

// DeleteTransaction implements repositories.TransactionsRepository.
func (i *ITransactionRepository) DeleteTransaction(id uuid.UUID) error {
	panic("unimplemented")
}

// GetTransactionByID implements repositories.TransactionsRepository.
func (i *ITransactionRepository) GetTransactionByID(id uuid.UUID) (repositories.Transactions, error) {
	panic("unimplemented")
}

// GetTransactionsByBondID implements repositories.TransactionsRepository.
func (i *ITransactionRepository) GetTransactionsByBondID(bondID uuid.UUID) ([]repositories.Transactions, error) {
	panic("unimplemented")
}

// GetTransactionsByUserID implements repositories.TransactionsRepository.
func (i *ITransactionRepository) GetTransactionsByUserID(userID string) ([]repositories.Transactions, error) {
	panic("unimplemented")
}

// UpdateTransaction implements repositories.TransactionsRepository.
func (i *ITransactionRepository) UpdateTransaction(transaction repositories.Transactions) error {
	panic("unimplemented")
}
