package datasources

import (
	"context"
	"database/sql"
	"os"

	"go.uber.org/zap"
)

func Database1Connection(logger *zap.SugaredLogger, ctx context.Context) *sql.DB {
	host := os.Getenv("DATABASE_POSTGRES_HOST")
	port := os.Getenv("DATABASE_POSTGRES_PORT")
	user := os.Getenv("DATABASE_POSTGRES_USER")
	password := os.Getenv("DATABASE_POSTGRES_PASSWORD")
	dbname := os.Getenv("DATABASE_POSTGRES_DBNAME")

	conn, err := Connection(ctx, logger, host, port, user, password, dbname, "disable")
	if err != nil {
		logger.Error("Error connecting to database")
	}
	return conn
}
