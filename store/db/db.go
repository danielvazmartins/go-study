package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	connStr := "user=postgres dbname=alura-store password='123456' sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
