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
	filter := bson.M{"email": email}
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)
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
	filterUser := bson.M{"_id": objID}
	err = userCollection.FindOne(context.Background(), filterUser).Decode(&user)
	if err != nil {
		return nil, err
	}

	if len(user.Friends) == 0 {
		return []models.User{}, nil
	}

	filterFriends := bson.M{"_id": bson.M{"$in": user.Friends}}

	// Find all friends by their ObjectIDs
	cursor, err := userCollection.Find(context.Background(), filterFriends)
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

// Get all Users from Database except the one with the given userId
func GetUsersFromDB(userId string) ([]models.User, error) {
	client := config.GetMongoDBClient()
	userCollection := client.Database("cosmic-gate-db").Collection("users")

	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{Key: "_id", Value: bson.D{{Key: "$ne", Value: objID}}},
	}

	cursor, err := userCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}
