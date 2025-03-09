package services

import (
	"context"
	"cosmic-gate-chat/config"
	"cosmic-gate-chat/models"
	"log"
)

// Saves Message into Collection
func SaveMessage(message models.Message) error {
	client := config.GetMongoDBClient()
	messagesColl := client.Database("cosmic-gate-db").Collection("messages")

	_, err := messagesColl.InsertOne(context.Background(), message)
	if err != nil {
		log.Println("SaveMessage:", err)
		return err
	}
	return nil
}
