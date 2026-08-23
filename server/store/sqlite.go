package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitaliseSqliteStore() *sql.DB {
	db, err := sql.Open("sqlite3", "./titly.db")

	if err != nil {
		panic(err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS url_mappings (
		id INTEGER PRIMARY KEY,
		long_url TEXT NOT NULL,
		short_url TEXT NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(schema)
	if err != nil {
		panic(err)
	}

	return db
}

