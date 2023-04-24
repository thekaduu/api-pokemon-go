package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func getConnectionString() string {
	return "user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable"
}

func Query(query string, args ...interface{}) *sql.Rows {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), getConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(query, args...)

	if err != nil {
		panic(err)
	}

	return rows
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), getConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db.QueryRow(query, args...)
}

func Count(query string) int {
	var count int

	db, err := sql.Open(os.Getenv("DB_DRIVER"), getConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.QueryRow(query).Scan(&count)

	if err != nil {
		panic(err)
	}

	return count
}
