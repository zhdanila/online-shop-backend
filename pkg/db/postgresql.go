package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"online-shop-backend/internal/config"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cnf *config.Config) (*sqlx.DB, error) {
	dbConfig := DBConfig{
		Host:     cnf.DBHost,
		Port:     cnf.DBPort,
		Username: cnf.DBUsername,
		DBName:   cnf.DBName,
		SSLMode:  cnf.DBSSLMode,
		Password: cnf.DBPassword,
	}

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.DBName, dbConfig.Password, dbConfig.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	zap.L().Info("DB connected")

	return db, nil
}
