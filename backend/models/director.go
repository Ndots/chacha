package models

import "database/sql"

type Director struct {
	ID         int    `json:"id"`
	BusinessID int    `json:"business_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

func CreateDirector(db *sql.DB, director *Director) error {
	query := `INSERT INTO directors (business_id, name, email) VALUES ($1, $2, $3) RETURNING id`
	return db.QueryRow(query, director.BusinessID, director.Name, director.Email).Scan(&director.ID)
}
