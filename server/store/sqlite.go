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
		id INTEGER PRIMARY KEY AUTOINCREMENT,
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

func ExistsInSqliteStore(db *sql.DB, longUrl string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM url_mappings WHERE long_url = ?);"

	err := db.QueryRow(query, longUrl).Scan(&exists)

	if err != nil {
		panic(err)
	}

	return exists
}

func GetShortUrlFromSqliteStore(db *sql.DB, longUrl string) string {
	var shortUrl string

	query := "SELECT short_url from url_mappings WHERE long_url = ?;"

	err := db.QueryRow(query, longUrl).Scan(&shortUrl)

	if err != nil {
		panic(err)
	}

	return shortUrl
}

func StoreMappingInSqliteStore(db *sql.DB, longUrl string, shortUrl string) {

}