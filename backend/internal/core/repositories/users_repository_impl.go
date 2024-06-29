package repositories

import (
	"database/sql"
	"io"
	"os"
	"technical-challenge/internal/core/domain/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IUsersRepository struct {
	Logger     *zap.SugaredLogger
	datasoruce *sql.DB
}

func NewUsersRepository(
	logger *zap.SugaredLogger,
	datasoruce *sql.DB,
) repositories.UsersRepository {
	logger.Info("UsersRepository created")

	IUsersRepository := &IUsersRepository{
		Logger:     logger,
		datasoruce: datasoruce,
	}

	//si os.Getenv("ENV")  es dev, test. local, etc
	if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "local" {
		logger.Info("Syncing database")
		IUsersRepository.SyncDatabase()
	}

	return IUsersRepository
}

// Función para sincronizar la base de datos ejecutando un archivo SQL.
func (i *IUsersRepository) SyncDatabase() {
	// Leer el contenido del archivo SQL
	var filepath string = "migrations/users.sql"

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
func (i *IUsersRepository) Exec(query string) {
	_, err := i.datasoruce.Exec(query)
	if err != nil {
		i.Logger.Error("Error executing query. ", err)
	}
}

// CreateUser implements repositories.UsersRepository.
func (i *IUsersRepository) CreateUser(user *repositories.User) error {
	panic("unimplemented")
}

// GetUserByEmail implements repositories.UsersRepository.
func (i *IUsersRepository) GetUserByEmail(email string) (*repositories.User, error) {
	panic("unimplemented")
}

// GetUserByID implements repositories.UsersRepository.
func (i *IUsersRepository) GetUserByID(id uuid.UUID) (*repositories.User, error) {
	panic("unimplemented")
}
