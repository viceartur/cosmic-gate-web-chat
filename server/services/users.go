package services

import (
	"context"
	"cosmic-gate-chat/config"
	"cosmic-gate-chat/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Get User from Database by Email
func GetUserFromDB(email string) (models.User, error) {
	client := config.GetMongoDBClient()
	userCollection := client.Database("cosmic-gate-db").Collection("users")

	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
