package seeds

import (
	"chacha/backend/database"
	"chacha/backend/models"
	"chacha/backend/utils"
	"log"
)

func Seed() {
	db := database.ConnectDB()
	defer db.Close()

	// Two normal users
	users := []models.User{
		{Name: "John Doe", Email: "john@example.com", Password: "password123", Address: "123 Main St", Role: "user"},
		{Name: "Jane Smith", Email: "jane@example.com", Password: "password123", Address: "456 Elm St", Role: "user"},
	}
	// One admin user
	admin := models.User{Name: "Admin User", Email: "admin@example.com", Password: "adminpass", Address: "789 Admin Ave", Role: "admin"}

	// One partner
	partner := models.Partner{Name: "Partner Authority", ApiKey: "partner-key", Contact: "partner@example.com"}

	// Seed normal users
	for _, u := range users {
		hashed, err := utils.HashPassword(u.Password)
		if err != nil {
			log.Println("Error hashing password for", u.Email, err)
			continue
		}
		u.Password = hashed
		if err := models.CreateUser(db, &u); err != nil {
			log.Println("Error creating user", u.Email, err)
		} else {
			log.Println("Created user", u.Email)
		}
	}

	// Seed admin user
	hashed, err := utils.HashPassword(admin.Password)
	if err != nil {
		log.Println("Error hashing admin password", err)
	}
	admin.Password = hashed
	if err := models.CreateUser(db, &admin); err != nil {
		log.Println("Error creating admin user", err)
	} else {
		log.Println("Created admin user", admin.Email)
	}

	// Seed partner
	if err := models.CreatePartner(db, &partner); err != nil {
		log.Println("Error creating partner", err)
	} else {
		log.Println("Created partner", partner.Name)
	}
}
