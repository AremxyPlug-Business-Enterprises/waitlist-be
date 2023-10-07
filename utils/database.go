package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var db *mongo.Database

func ConnectDatabase() *mongo.Database {
	uri := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	log.Println("DB_URL:", uri)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(dbName)
	return db
}
