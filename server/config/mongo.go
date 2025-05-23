package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// MongoDB initialization
func InitMongoDB() {
	uri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	MongoClient = client
	log.Println("MongoDB connected")
}

// Returns the pointer to Mongo Client Connection
func GetMongoDBClient() *mongo.Client {
	return MongoClient
}
