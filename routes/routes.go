package routes

import (
	"github.com/gorilla/mux"
	"github.com/waitlist-be/controllers"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/api/addWaitlist", controllers.AddToWaitlist).Methods("POST")
	router.HandleFunc("/api/getWaitlist", controllers.GetWaitlist).Methods("GET")
}
