package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func setupDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://user:0000@localhost/mydatabase?sslmode=disable")
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
		return nil, err
	}

	return db, nil
}
