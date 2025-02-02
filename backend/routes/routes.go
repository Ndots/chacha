package routes

import (
	"chacha/backend/handlers"
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// Auth endpoints
	router.HandleFunc("/api/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/login", handlers.LoginUser).Methods("POST")

	// Business endpoints
	router.HandleFunc("/api/business", handlers.RegisterBusiness).Methods("POST")

	// Application endpoints
	router.HandleFunc("/api/application", handlers.UpdateApplication).Methods("PUT")

	// Dashboard endpoint (for admin & user)
	router.HandleFunc("/api/dashboard", handlers.GetDashboardStats).Methods("GET")

	// Partner endpoint
	router.HandleFunc("/api/partner/application", handlers.PartnerApproveReject).Methods("PUT")

	return router
}
