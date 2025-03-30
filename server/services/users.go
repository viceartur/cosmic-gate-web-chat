package services

import (
	"context"
	"cosmic-gate-chat/config"
	"cosmic-gate-chat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Get User Friends by their ObjectID
func GetUserFriendsFromDB(userID string) ([]models.User, error) {
	client := config.GetMongoDBClient()
	userCollection := client.Database("cosmic-gate-db").Collection("users")

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	// Find the user and get their friends list
	var user models.User
	err = userCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	if len(user.Friends) == 0 {
		return []models.User{}, nil
	}

	// Find all friends by their ObjectIDs
	cursor, err := userCollection.Find(context.Background(), bson.M{"_id": bson.M{"$in": user.Friends}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var friends []models.User
	if err := cursor.All(context.Background(), &friends); err != nil {
		return nil, err
	}

	return friends, nil
}
