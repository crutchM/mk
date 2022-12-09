package repositories

import "github.com/jmoiron/sqlx"

func NewDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "notes.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
