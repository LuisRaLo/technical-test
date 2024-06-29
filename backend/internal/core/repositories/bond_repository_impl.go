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
) repositories.BoundRepository {
	logger.Info("BoundRepository created")

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

// CreateBound implements repositories.BoundRepository.
func (i *IBondRepository) CreateBound(bound repositories.Bound) error {
	panic("unimplemented")
}

// DeleteBound implements repositories.BoundRepository.
func (i *IBondRepository) DeleteBound(bound repositories.Bound) error {
	panic("unimplemented")
}

// GetAllBounds implements repositories.BoundRepository.
func (i *IBondRepository) GetAllBounds() ([]repositories.Bound, error) {
	panic("unimplemented")
}

// GetBoundByID implements repositories.BoundRepository.
func (i *IBondRepository) GetBoundByID(id uuid.UUID) (repositories.Bound, error) {
	panic("unimplemented")
}

// UpdateBound implements repositories.BoundRepository.
func (i *IBondRepository) UpdateBound(bound repositories.Bound) error {
	panic("unimplemented")
}
