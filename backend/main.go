// entry point
package main

import (
	"log"
	"net/http"

	// "chacha/backend/seeds" // uncomment this line to seed the database
	// For PostgreSQL driver
	"chacha/backend/database"
	"chacha/backend/routes"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	// Initialize database connection
	db := database.ConnectDB()
	defer db.Close()

	// Uncomment this line to seed the database
	//  seeds.Seed()

	// Run database migrations (if any)
	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize routes
	router := routes.InitializeRoutes()

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	// Start the server with CORS-enabled handler
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("Server error:", err)
	}
}
