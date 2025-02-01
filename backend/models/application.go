package models

import "database/sql"

type Application struct {
	ID         int    `json:"id"`
	BusinessID int    `json:"business_id"`
	Status     string `json:"status"`
	Reason     string `json:"reason,omitempty"`
	PartnerID  int    `json:"partner_id,omitempty"`
}

func CreateApplication(db *sql.DB, app *Application) error {
	query := `INSERT INTO applications (business_id, status, reason, partner_id) VALUES ($1, $2, $3, $4) RETURNING id`
	return db.QueryRow(query, app.BusinessID, app.Status, app.Reason, app.PartnerID).Scan(&app.ID)
}

func UpdateApplicationStatus(db *sql.DB, applicationID int, status string, reason string) error {
	query := `UPDATE applications SET status = $1, reason = $2 WHERE id = $3`
	_, err := db.Exec(query, status, reason, applicationID)
	return err
}

func GetApplicationsByUser(db *sql.DB, userID int) ([]Application, error) {
	query := `SELECT a.id, a.business_id, a.status, a.reason, a.partner_id
	          FROM applications a 
	          JOIN businesses b ON a.business_id = b.id 
	          WHERE b.user_id = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var apps []Application
	for rows.Next() {
		var a Application
		if err := rows.Scan(&a.ID, &a.BusinessID, &a.Status, &a.Reason, &a.PartnerID); err != nil {
			return nil, err
		}
		apps = append(apps, a)
	}
	return apps, nil
}
