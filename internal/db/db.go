package db

import (
	"database/sql"
	"errors"
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
				login CHARACTER VARYING(100) NOT NULL UNIQUE,
				password CHARACTER VARYING(100) NOT NULL,
				balance INTEGER NOT NULL DEFAULT 0 CHECK(balance >= 0),
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			)
		`,
		`
			CREATE TABLE IF NOT EXISTS orders
			(
				id SERIAL PRIMARY KEY,
				uid INTEGER REFERENCES users (id),
				number INTEGER NOT NULL UNIQUE,
				bonuses INTEGER,
				status VARYING(10),
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
		//насколько правильное такое написание? С одной стороны, написано в 1 строку, а с другой, не совсем явно, что я тут ошибку возвращаю (или nil)
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
