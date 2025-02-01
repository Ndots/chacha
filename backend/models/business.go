package models

import "database/sql"

type Business struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	Type          string   `json:"type"`
	ProposedNames []string `json:"proposed_names"`
	Status        string   `json:"status"`
	UserID        int      `json:"user_id"`
}

func CreateBusiness(db *sql.DB, business *Business) error {
	query := `INSERT INTO businesses (name, address, type, proposed_names, status, user_id) VALUES ($1, $2, $3, $4, 'Pending', $5) RETURNING id`
	return db.QueryRow(query, business.Name, business.Address, business.Type, business.ProposedNames, business.UserID).Scan(&business.ID)
}

func GetBusinessesByUser(db *sql.DB, userID int) ([]Business, error) {
	query := `SELECT id, name, address, type, proposed_names, status, user_id FROM businesses WHERE user_id = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var businesses []Business
	for rows.Next() {
		var b Business
		if err := rows.Scan(&b.ID, &b.Name, &b.Address, &b.Type, &b.ProposedNames, &b.Status, &b.UserID); err != nil {
			return nil, err
		}
		businesses = append(businesses, b)
	}
	return businesses, nil
}
