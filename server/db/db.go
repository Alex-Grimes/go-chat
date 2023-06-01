package db

import (
	"database/sql"
)

type Database struct {
	db *sql.DB
}

func newDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgressql://root:password@localhost:5433/go-chat?sslmode=disable")

	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
