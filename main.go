package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/waitlist-be/middleware"
	"github.com/waitlist-be/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	// Add the CORS middleware to your router
	router.Use(middleware.CorsMiddleware)

	routes.SetupRoutes(router)

	port := ":" + os.Getenv("PORT")
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
