package datasources

import (
	"fmt"
	"strconv"

	"database/sql"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func Connection(
	logger *zap.SugaredLogger,
	host string,
	port string,
	user string,
	password string,
	dbname string,
	sslmode string,
) (*sql.DB, error) {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		logger.Errorln("error converting port to integer: %v", err)
		return nil, err
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, portInt, user, password, dbname, sslmode))

	if err != nil {
		logger.Errorln("error connecting to database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logger.Errorln("Database connection failed: %v", err)
		return nil, err
	}
	logger.Infoln("Database connection successful")
	return db, nil

}

func CloseDB(db *sql.DB, logger *zap.SugaredLogger) {
	db.Close()
	logger.Infoln("Database connection closed")
}
