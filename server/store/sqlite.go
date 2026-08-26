package store

import (
	"context"
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

type UrlMapping struct {
	ID        int64  `json:"id"`
	LongUrl   string `json:"long_url"`
	ShortUrl  string `json:"short_url"`
	CreatedAt string `json:"created_at"`
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

func StoreMappingInSqliteStore(db *sql.DB, longUrl string, shortUrl string) int64 {
	query := "INSERT INTO url_mappings (long_url, short_url, created_at) VALUES (?, ?, CURRENT_TIMESTAMP);"

	id, err := db.Exec(query, longUrl, shortUrl)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := id.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId
}

func GetLongUrlFromSqliteStore(ctx context.Context, db *sql.DB, shortUrl string) (string, error) {
	var longUrl string
	query := "SELECT long_url FROM url_mappings WHERE short_url = ?;"

	err := db.QueryRowContext(ctx, query, shortUrl).Scan(&longUrl)
	if err != nil {
		return "", err
	}

	return longUrl, nil
}