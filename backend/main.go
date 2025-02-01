package main

import (
	"log"
	"net/http"

	"chacha/backend/database"
	"chacha/backend/routes"

	// "chacha/backend/seeds" // uncomment this line to seed the database
	// For PostgreSQL driver
	_ "github.com/lib/pq"
)

func main() {
	// Initialize database connection
	db := database.ConnectDB()
	defer db.Close()

	// Run database migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize routes
	router := routes.InitializeRoutes()

	// Start the server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
