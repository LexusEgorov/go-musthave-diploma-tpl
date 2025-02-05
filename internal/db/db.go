package db

import (
	"database/sql"
	"errors"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	DB *sql.DB
}

var db *DB

func (d *DB) createTables() error {
	if d.DB == nil {
		return errors.New("db isn't found")
	}

	createQueries := [2]string{
		`
			CREATE TABLE IF NOT EXISTS users 
			(
				id SERIAL PRIMARY KEY,
				login VARCHAR(100) NOT NULL UNIQUE,
				password VARCHAR(100) NOT NULL,
				balance INTEGER NOT NULL DEFAULT 0 CHECK(balance >= 0),
				balance_wd INTEGER NOT NULL DEFAULT 0 CHECK(balance_wd >= 0),
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS orders
			(
				id SERIAL PRIMARY KEY,
				uid INTEGER REFERENCES users (id),
				number VARCHAR(15) NOT NULL UNIQUE,
				bonuses INTEGER,
				status VARCHAR(10),
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			)
		`,
	}

	for _, q := range createQueries {
		_, err := d.DB.Exec(q)

		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DB) connect(connection string, isCreated bool) error {
	var err error
	d.DB, err = sql.Open("pgx", connection)

	if err != nil {
		return err
	}

	err = d.DB.Ping()

	if err != nil {
		return err
	}

	if !isCreated {
		return d.createTables()
	}

	return nil
}

func NewDB(connection string, isCreated bool) *DB {
	if db != nil {
		return db
	}

	db = &DB{}

	db.connect(connection, isCreated)

	return db
}
