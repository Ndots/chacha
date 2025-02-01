package models

import "database/sql"

type Partner struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	ApiKey  string `json:"api_key"`
	Contact string `json:"contact"`
}

func CreatePartner(db *sql.DB, partner *Partner) error {
	query := `INSERT INTO partners (name, api_key, contact) VALUES ($1, $2, $3) RETURNING id`
	return db.QueryRow(query, partner.Name, partner.ApiKey, partner.Contact).Scan(&partner.ID)
}

func GetPartnerByApiKey(db *sql.DB, apiKey string) (*Partner, error) {
	partner := &Partner{}
	query := `SELECT id, name, api_key, contact FROM partners WHERE api_key = $1`
	err := db.QueryRow(query, apiKey).Scan(&partner.ID, &partner.Name, &partner.ApiKey, &partner.Contact)
	return partner, err
}
