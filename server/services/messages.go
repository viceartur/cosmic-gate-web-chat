package services

import (
	"context"
	"cosmic-gate-chat/config"
	"cosmic-gate-chat/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

// Get all Messages between two users
func GetMessages(senderId string, recipientId string) ([]models.Message, error) {
	client := config.GetMongoDBClient()
	messageCollection := client.Database("cosmic-gate-db").Collection("messages")

	// Find all messages for users
	cursor, err := messageCollection.Find(context.Background(), bson.M{
		"$or": []bson.M{
			{"senderId": senderId, "recipientId": recipientId},
			{"senderId": recipientId, "recipientId": senderId},
		},
	})
	if err != nil {
		return []models.Message{}, err
	}

	// Decode documents into results
	var messages []models.Message
	err = cursor.All(context.Background(), &messages)
	if err != nil {
		return []models.Message{}, err
	}

	return messages, nil
}
