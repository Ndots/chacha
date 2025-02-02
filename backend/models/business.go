package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Director struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Position string `json:"position"`
}

type Business struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Address   string     `json:"address"`
	Type      string     `json:"type"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Directors []Director `json:"directors"`
	Status    string     `json:"status"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func CreateBusiness(db *sql.DB, business *Business) error {
	directorsJSON, err := json.Marshal(business.Directors)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO businesses (name, address, type, email, phone, directors, status, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
		RETURNING id, created_at, updated_at`

	return db.QueryRow(
		query,
		business.Name,
		business.Address,
		business.Type,
		business.Email,
		business.Phone,
		directorsJSON,
		business.Status,
		business.UserID,
	).Scan(&business.ID, &business.CreatedAt, &business.UpdatedAt)
}

func GetBusinessesByUser(db *sql.DB, userID int) ([]Business, error) {
	query := `SELECT id, name, address, type, email, phone, directors, status, created_at, updated_at 
              FROM businesses WHERE user_id = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var businesses []Business
	for rows.Next() {
		var b Business
		var directorsJSON []byte
		if err := rows.Scan(&b.ID, &b.Name, &b.Address, &b.Type, &b.Email,
			&b.Phone, &directorsJSON, &b.Status, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(directorsJSON, &b.Directors); err != nil {
			return nil, err
		}
		businesses = append(businesses, b)
	}
	return businesses, nil
}
