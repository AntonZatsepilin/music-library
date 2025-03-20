package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)


type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLmode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	logrus.Info("Connecting to the database")
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLmode))
	if err != nil {
		return nil, err
	}

	logrus.Info("Checking the connection to the database")
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}