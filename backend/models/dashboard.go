package models

import (
	"database/sql"
	"log"
)

type DashboardStats struct {
	TotalUsers           int `json:"total_users"`
	TotalBusinesses      int `json:"total_businesses"`
	PendingApplications  int `json:"pending_applications"`
	ApprovedApplications int `json:"approved_applications"`
	RejectedApplications int `json:"rejected_applications"`
}

func GetDashboardStats(db *sql.DB) (*DashboardStats, error) {
	stats := &DashboardStats{}
	query := `SELECT COUNT(*) FROM users`
	if err := db.QueryRow(query).Scan(&stats.TotalUsers); err != nil {
		log.Println("Error fetching total users:", err)
	}
	query = `SELECT COUNT(*) FROM businesses`
	if err := db.QueryRow(query).Scan(&stats.TotalBusinesses); err != nil {
		log.Println("Error fetching total businesses:", err)
	}
	query = `SELECT COUNT(*) FROM applications WHERE status = 'Pending'`
	if err := db.QueryRow(query).Scan(&stats.PendingApplications); err != nil {
		log.Println("Error fetching pending applications:", err)
	}
	query = `SELECT COUNT(*) FROM applications WHERE status = 'Approved'`
	if err := db.QueryRow(query).Scan(&stats.ApprovedApplications); err != nil {
		log.Println("Error fetching approved applications:", err)
	}
	query = `SELECT COUNT(*) FROM applications WHERE status = 'Rejected'`
	if err := db.QueryRow(query).Scan(&stats.RejectedApplications); err != nil {
		log.Println("Error fetching rejected applications:", err)
	}
	return stats, nil
}
