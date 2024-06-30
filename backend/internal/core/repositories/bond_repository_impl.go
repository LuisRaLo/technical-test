package repositories

import (
	"database/sql"
	"io"
	"os"
	"technical-challenge/internal/core/domain/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IBondRepository struct {
	Logger     *zap.SugaredLogger
	datasoruce *sql.DB
}

func NewBondRepository(
	logger *zap.SugaredLogger,
	datasoruce *sql.DB,
) repositories.BondRepository {
	logger.Info("BondRepository created")

	IBondRepository := &IBondRepository{
		Logger:     logger,
		datasoruce: datasoruce,
	}

	if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "local" {
		logger.Info("Syncing database")
		IBondRepository.SyncDatabase()
	}

	return IBondRepository
}

// Función para sincronizar la base de datos ejecutando un archivo SQL.
func (i *IBondRepository) SyncDatabase() {
	var filepath string = "test/migrations/bonds.sql"

	file, err := os.Open(filepath)
	if err != nil {
		i.Logger.Error("Error opening file")
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		i.Logger.Error("Error reading file")
	}

	i.Logger.Info(string(content))

	_, err = i.datasoruce.Exec(string(content))
	if err != nil {
		i.Logger.Error("Error executing SQL file. ", err)
	}

	i.Logger.Info("Database synced")
}

// Exec
func (i *IBondRepository) Exec(query string) {
	_, err := i.datasoruce.Exec(query)
	if err != nil {
		i.Logger.Error("Error executing query. ", err)
	}
}

// CreateBond implements repositories.BondRepository.
func (i *IBondRepository) CreateBond(bond repositories.Bond) int {
	query := `INSERT INTO fpd.bonds (id, user_id, name, quantity, price, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	r, err := i.datasoruce.Exec(query, bond.ID, bond.UserID, bond.Name, bond.Quantity, bond.Price, bond.CreatedAt, bond.UpdatedAt, bond.DeleteAt)
	if err != nil {
		i.Logger.Error("Error creating bond. ", err)
		return 0
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		i.Logger.Error("Error getting rows affected. ", err)
		return 0
	}
	return int(rowsAffected)

}

// DeleteBond implements repositories.BondRepository.
func (i *IBondRepository) DeleteBond(bond repositories.Bond) int {
	panic("unimplemented")
}

// GetAllBonds implements repositories.BondRepository.
func (i *IBondRepository) GetAllBonds() ([]repositories.Bond, error) {
	panic("unimplemented")
}

// GetBondByID implements repositories.BondRepository.
func (i *IBondRepository) GetBondByID(id uuid.UUID) (repositories.Bond, error) {
	sqlStatement := `SELECT * FROM fpd.bonds WHERE id = $1`
	row := i.datasoruce.QueryRow(sqlStatement, id)

	var bond repositories.Bond
	err := row.Scan(&bond.ID, &bond.UserID, &bond.Name, &bond.Quantity, &bond.Price, &bond.CreatedAt, &bond.UpdatedAt, &bond.DeleteAt)
	if err != nil {
		i.Logger.Error("Error getting bond by ID. ", err)
		return repositories.Bond{}, err
	}

	return bond, nil
}

func (i *IBondRepository) IsExistBond(name string, price float64, quantity int) bool {
	query := `SELECT id FROM fpd.bonds WHERE EXISTS (SELECT * FROM fpd.bonds WHERE name = $1 AND price = $2 AND quantity = $3)`
	row := i.datasoruce.QueryRow(query, name, price, quantity)

	var id uuid.UUID
	err := row.Scan(&id)
	return err == nil
}

// UpdateBond implements repositories.BondRepository.
func (i *IBondRepository) UpdateBond(bond repositories.Bond) int {
	panic("unimplemented")
}
