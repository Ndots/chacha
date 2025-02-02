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
			name VARCHAR(255) NOT NULL,
			address TEXT NOT NULL,
			type VARCHAR(100) NOT NULL,
			email VARCHAR(255) NOT NULL,
			phone VARCHAR(20) NOT NULL,
			directors JSONB NOT NULL DEFAULT '[]'::jsonb,
			status VARCHAR(50) NOT NULL DEFAULT 'pending',
			user_id INTEGER REFERENCES users(id),
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,
		// Directors table --- optional
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
