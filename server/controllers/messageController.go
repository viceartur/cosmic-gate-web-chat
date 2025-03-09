package controllers

import (
	"context"
	"cosmic-gate-chat/config"
	"cosmic-gate-chat/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

// Get Messages from DB for all SenderID-RecipientID combinations
func GetMessages(w http.ResponseWriter, r *http.Request) {
	client := config.GetMongoDBClient()
	messageCollection := client.Database("cosmic-gate-db").Collection("messages")

	senderId, _ := strconv.Atoi(r.URL.Query().Get("senderId"))
	recipientId, _ := strconv.Atoi(r.URL.Query().Get("recipientId"))

	// Find all messages for users
	cursor, err := messageCollection.Find(context.Background(), bson.M{
		"$or": []bson.M{
			{"senderId": senderId, "recipientId": recipientId},
			{"senderId": recipientId, "recipientId": senderId},
		},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching messages: %v", err), http.StatusInternalServerError)
		return
	}

	// Decode documents into results
	var messages []models.Message
	err = cursor.All(context.Background(), &messages)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding messages: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
