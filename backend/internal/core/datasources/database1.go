package datasources

import (
	"database/sql"
	"os"

	"go.uber.org/zap"
)

func Database1Connection(logger *zap.SugaredLogger) *sql.DB {
	host := os.Getenv("DATABASE_POSTGRES_HOST")
	port := os.Getenv("DATABASE_POSTGRES_PORT")
	user := os.Getenv("DATABASE_POSTGRES_USER")
	password := os.Getenv("DATABASE_POSTGRES_PASSWORD")
	dbname := os.Getenv("DATABASE_POSTGRES_DBNAME")

	conn, err := Connection(logger, host, port, user, password, dbname, "disable")
	if err != nil {
		logger.Error("Error connecting to database")
	}
	return conn
}
