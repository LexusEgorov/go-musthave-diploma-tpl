package db

import "database/sql"

var db *sql.DB

type DB struct {
	db *sql.DB
}

func (d *DB) createTables() error {

	return nil
}

func (d *DB) connect(connection string, isCreated bool) error {
	if db != nil {
		d.db = db
		return nil
	}

	var err error
	d.db, err = sql.Open("pgx", connection)

	if err != nil {
		return err
	}

	err = d.db.Ping()

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
	db := DB{}

	db.connect(connection, isCreated)

	return &db
}
