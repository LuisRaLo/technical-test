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
	var filepath string = "test/migrations/transactions.sql"

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
func (i *ITransactionRepository) CreateTransaction(transaction repositories.Transactions) (int64, error) {
	stmt, err := i.datasoruce.Prepare("INSERT INTO fpd.transactions(id, user_id, bond_id, status, quantity, price, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		i.Logger.Error("Error preparing statement. ", err)
		return 0, err
	}

	res, err := stmt.Exec(transaction.ID, transaction.UserID, transaction.BondID, transaction.Status, transaction.Quantity, transaction.Price, transaction.CreatedAt, transaction.UpdatedAt, transaction.DeletedAt)
	if err != nil {
		i.Logger.Error("Error executing statement. ", err)
		return 0, err
	}

	id, err := res.RowsAffected()
	if err != nil {
		i.Logger.Error("Error getting last insert ID. ", err)
		return 0, err
	}

	return id, nil
}

// DeleteTransaction implements repositories.TransactionsRepository.
func (i *ITransactionRepository) DeleteTransaction(id uuid.UUID) error {
	panic("unimplemented")
}

// GetTransactionByID implements repositories.TransactionsRepository.
func (i *ITransactionRepository) GetTransactionByID(id uuid.UUID) (repositories.Transactions, error) {
	stmt, err := i.datasoruce.Prepare("SELECT * FROM fpd.transactions WHERE id = $1")
	if err != nil {
		i.Logger.Error("Error preparing statement. ", err)
		return repositories.Transactions{}, err
	}

	row := stmt.QueryRow(id)
	var transaction repositories.Transactions
	err = row.Scan(&transaction.ID, &transaction.UserID, &transaction.BondID, &transaction.Status, &transaction.Quantity, &transaction.Price, &transaction.CreatedAt, &transaction.UpdatedAt, &transaction.DeletedAt)
	if err != nil {
		i.Logger.Error("Error scanning row. ", err)
		return repositories.Transactions{}, err
	}

	return transaction, nil
}

// GetTransactionsByBondID implements repositories.TransactionsRepository.
func (i *ITransactionRepository) GetTransactionsByBondID(bondID uuid.UUID) ([]repositories.Transactions, error) {
	panic("unimplemented")
}

// GetTransactionsByUserID implements repositories.TransactionsRepository.
func (i *ITransactionRepository) GetTransactionsByUserID(userID string) ([]repositories.Transactions, error) {
	stmt, err := i.datasoruce.Prepare("SELECT * FROM fpd.transactions WHERE user_id = $1")
	if err != nil {
		i.Logger.Error("Error preparing statement. ", err)
		return nil, err
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		i.Logger.Error("Error getting transactions by user ID. ", err)
		return nil, err
	}

	var transactions []repositories.Transactions
	for rows.Next() {
		var transaction repositories.Transactions
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.BondID, &transaction.Status, &transaction.Quantity, &transaction.Price, &transaction.CreatedAt, &transaction.UpdatedAt, &transaction.DeletedAt)
		if err != nil {
			i.Logger.Error("Error scanning row. ", err)
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (i *ITransactionRepository) GetTransactionsByUserIDAndStatus(userID string, status string) ([]repositories.Transactions, error) {
	sql := `
		SELECT *
		FROM fpd.transactions AS t
		WHERE t.user_id = $1
			AND t.status = $2
	`

	stmt, err := i.datasoruce.Prepare(sql)
	if err != nil {
		i.Logger.Error("Error preparing statement. ", err)
		return nil, err
	}

	rows, err := stmt.Query(userID, status)
	if err != nil {
		i.Logger.Error("Error getting transactions by user ID and status. ", err)
		return nil, err
	}

	var transactions []repositories.Transactions
	for rows.Next() {
		var transaction repositories.Transactions
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.BondID, &transaction.Status, &transaction.Quantity, &transaction.Price, &transaction.CreatedAt, &transaction.UpdatedAt, &transaction.DeletedAt)
		if err != nil {
			i.Logger.Error("Error scanning row. ", err)
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// ahora quiero toso los transactions que tengan el status de BOUGHT pero que no sean del usuario que esta haciendo la peticion
func (i *ITransactionRepository) GetTransactionsByStatusAndNotUserID(status string, userID string) ([]repositories.Transactions, error) {
	sql := `
		SELECT *
		FROM fpd.transactions AS t
		WHERE t.status = $1
			AND t.user_id != $2
	`

	i.Logger.Infow("SQL: ", "sql", sql)

	stmt, err := i.datasoruce.Prepare(sql)
	if err != nil {
		i.Logger.Error("Error preparing statement. ", err)
		return nil, err
	}

	rows, err := stmt.Query(status, userID)
	if err != nil {
		i.Logger.Error("Error getting transactions by status and not user ID. ", err)
		return nil, err
	}

	var transactions []repositories.Transactions
	for rows.Next() {
		var transaction repositories.Transactions
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.BondID, &transaction.Status, &transaction.Quantity, &transaction.Price, &transaction.CreatedAt, &transaction.UpdatedAt, &transaction.DeletedAt)
		if err != nil {
			i.Logger.Error("Error scanning row. ", err)
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// UpdateTransaction implements repositories.TransactionsRepository.
func (i *ITransactionRepository) UpdateTransaction(transaction repositories.Transactions) error {
	panic("unimplemented")
}
