package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nest-hackathon/server/config"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(cfg *config.Config) (*Database, error) {

	db, err := sql.Open("sqlite3", "./db/dev.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}
