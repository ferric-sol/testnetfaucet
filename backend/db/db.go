package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) (*Database, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	// Verify connection is working
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}
