package database

import (
	"database/sql"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	dbSource := os.Getenv("DATABASE_URL")
	if dbSource == "" {
		//dbSource = "user=chacha_user password=password dbname=chacha sslmode=disable" OR
		//dbSource = "user=postgres password=password dbname=chacha sslmode=disable"
		dbSource = "user=chacha_user password=password dbname=chacha sslmode=disable"
	}
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}
	log.Println("Connected to PostgreSQL database")
	return db
}
