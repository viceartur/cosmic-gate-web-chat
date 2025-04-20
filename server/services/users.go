package services

import (
	"context"
	"cosmic-gate-chat/config"
	"cosmic-gate-chat/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get User from Database by Email
func GetUserByEmail(email string) (models.User, error) {
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

// Get User from Database by ID
func GetUserById(userId string) (models.User, error) {
	client := config.GetMongoDBClient()
	userCollection := client.Database("cosmic-gate-db").Collection("users")

	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	filter := bson.M{"_id": userObjId}
	err = userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Get User Friends by their ObjectID
func GetUserFriends(userID string) ([]models.User, error) {
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
func GetUsers(userId string) ([]models.User, error) {
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

// Send a Friend Request to a User
func SendFriendRequest(userId string, friendId string) error {
	if userId == "" || friendId == "" {
		return errors.New("User ID and Friend ID cannot be empty")
	}

	client := config.GetMongoDBClient()
	userCollection := client.Database("cosmic-gate-db").Collection("users")

	ctx := context.Background()
	defer ctx.Done()

	// Convert string IDs to ObjectIDs
	userObjId, err := primitive.ObjectIDFromHex(userId)
	friendObjId, err := primitive.ObjectIDFromHex(friendId)
	if err != nil {
		return err
	}

	// Check if friendId user exists
	filter := bson.M{"_id": friendObjId}
	update := bson.M{
		"$addToSet": bson.M{
			"friendRequests": userObjId, // Adds userId if not already present
		},
	}

	updatedResult, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if updatedResult.MatchedCount == 0 {
		return errors.New("Friend user not found")
	}

	return nil
}

// Get User Friend Requests by their ObjectID
func GetUserFriendRequests(userId string) ([]models.User, error) {
	client := config.GetMongoDBClient()
	userCollection := client.Database("cosmic-gate-db").Collection("users")

	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	// Find the user and get their friend requests
	var user models.User
	filterUser := bson.M{"_id": objID}
	err = userCollection.FindOne(context.Background(), filterUser).Decode(&user)
	if err != nil {
		return nil, err
	}

	if len(user.FriendRequests) == 0 {
		return []models.User{}, nil
	}

	filterFriendRequests := bson.M{"_id": bson.M{"$in": user.FriendRequests}}

	// Find all Users from Friend Requests by their ObjectIDs
	cursor, err := userCollection.Find(context.Background(), filterFriendRequests)
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

// Accept a friend request between two users
func AcceptFriendRequest(userId string, friendId string) error {
	if userId == "" || friendId == "" {
		return errors.New("User ID and Friend ID cannot be empty")
	}

	client := config.GetMongoDBClient()
	userCollection := client.Database("cosmic-gate-db").Collection("users")

	ctx := context.Background()
	defer ctx.Done()

	// Convert to ObjectIDs
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	friendObjId, err := primitive.ObjectIDFromHex(friendId)
	if err != nil {
		return err
	}

	// Update user: remove from friendRequests, add to friends
	userUpdate := bson.M{
		"$pull": bson.M{
			"friendRequests": friendObjId,
		},
		"$addToSet": bson.M{
			"friends": friendObjId,
		},
	}

	userResult, err := userCollection.UpdateByID(ctx, userObjId, userUpdate)
	if err != nil {
		return err
	}
	if userResult.MatchedCount == 0 {
		return errors.New("User not found")
	}

	// Update friend: add user to their friends
	friendUpdate := bson.M{
		"$addToSet": bson.M{
			"friends": userObjId,
		},
	}

	friendResult, err := userCollection.UpdateByID(ctx, friendObjId, friendUpdate)
	if err != nil {
		return err
	}
	if friendResult.MatchedCount == 0 {
		return errors.New("Friend user not found")
	}

	return nil
}

// Remove a friend request without adding to friends
func DeclineFriendRequest(userId string, friendId string) error {
	if userId == "" || friendId == "" {
		return errors.New("User ID and Friend ID cannot be empty")
	}

	client := config.GetMongoDBClient()
	userCollection := client.Database("cosmic-gate-db").Collection("users")

	ctx := context.Background()
	defer ctx.Done()

	// Convert to ObjectIDs
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	friendObjId, err := primitive.ObjectIDFromHex(friendId)
	if err != nil {
		return err
	}

	// Remove friendId from userId's friendRequests
	update := bson.M{
		"$pull": bson.M{
			"friendRequests": friendObjId,
		},
	}

	result, err := userCollection.UpdateByID(ctx, userObjId, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("User not found")
	}

	return nil
}
