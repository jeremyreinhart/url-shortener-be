package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB


func ConnectDatabase() error {
	psqlInfo := os.Getenv("DATABASE_URL")
	if psqlInfo == "" {
	log.Fatal("DATABASE_URL not set")
	}
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return  err
	}
	DB = db
	return nil
}