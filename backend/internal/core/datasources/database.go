package datasources

import (
	"context"
	"fmt"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func Connection(
	ctx context.Context,
	logger *zap.SugaredLogger,
	host string,
	port string,
	user string,
	password string,
	dbname string,
	sslmode string,
) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s connect_timeout=10 sslmode=%s",
		host, port, dbname, user, password, sslmode)

	logger.Infoln("Connecting to database with connection string: %s", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	logger.Infoln("Database connection successful")
	return db, nil

}

func CloseDB(db *sql.DB, logger *zap.SugaredLogger) {
	db.Close()
	logger.Infoln("Database connection closed")
}
