package repositories

import (
	"database/sql"
	"io"
	"os"
	"technical-challenge/internal/core/domain/repositories"

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
	var filepath string = "test/migrations/users.sql"

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
func (i *IUsersRepository) CreateUser(user *repositories.User) int {

	query := `INSERT INTO users.users (id, email, name, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6);`
	result, err := i.datasoruce.Exec(query, user.ID, user.Email, user.Name, user.CreatedAt, user.UpdatedAt, user.DeleteAt)
	if err != nil {
		i.Logger.Error("Error executing query. ", err)
		return 0
	}
	r, err := result.RowsAffected()
	if err != nil {
		i.Logger.Error("Error executing query. ", err)
	}
	return int(r)

}

// GetUserByID implements repositories.UsersRepository.
func (i *IUsersRepository) GetUserByID(id string) (repositories.User, error) {
	query := "SELECT u.id, u.email, u.name, u.created_at, u.updated_at, u.deleted_at FROM users.users AS u WHERE id = $1;"
	rows, err := i.datasoruce.Query(query, id)
	if err != nil {
		i.Logger.Error("Error executing query. ", err)
	}
	defer rows.Close()

	var user repositories.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt, &user.DeleteAt)
		if err != nil {
			i.Logger.Error("Error scanning row. ", err)
		}
	}

	return user, nil

}

// isEmailAlreadyInUse implements repositories.UsersRepository.
func (i *IUsersRepository) IsEmailAlreadyInUse(email string) bool {
	query := "SELECT u.id FROM users.users AS u WHERE EXISTS (SELECT id FROM users.users WHERE email = $1);"
	rows, err := i.datasoruce.Query(query, email)
	if err != nil {
		i.Logger.Error("Error executing query. ", err)
	}
	defer rows.Close()
	return rows.Next()

}
