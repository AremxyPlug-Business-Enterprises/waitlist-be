package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"waitlist-be/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	routes.SetupRoutes(r) // Pass the router to SetupRoutes

	port := ":" + os.Getenv("PORT")
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
