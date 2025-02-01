package database

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) error {
	migrations := []string{
		// Users table
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			address TEXT,
			role TEXT NOT NULL DEFAULT 'user',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		// Businesses table
		`CREATE TABLE IF NOT EXISTS businesses (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			address TEXT,
			type TEXT,
			proposed_names TEXT[],
			status TEXT NOT NULL DEFAULT 'Pending',
			user_id INT REFERENCES users(id),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		// Directors table
		`CREATE TABLE IF NOT EXISTS directors (
			id SERIAL PRIMARY KEY,
			business_id INT REFERENCES businesses(id),
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL
		);`,
		// Partners table
		`CREATE TABLE IF NOT EXISTS partners (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			api_key TEXT UNIQUE NOT NULL,
			contact TEXT NOT NULL
		);`,
		// Applications table
		`CREATE TABLE IF NOT EXISTS applications (
			id SERIAL PRIMARY KEY,
			business_id INT REFERENCES businesses(id),
			status TEXT NOT NULL,
			reason TEXT,
			partner_id INT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
	}

	for _, query := range migrations {
		if _, err := db.Exec(query); err != nil {
			return err
		}
	}
	log.Println("Migrations completed successfully")
	return nil
}
