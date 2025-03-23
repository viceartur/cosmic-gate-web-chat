package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Username       string               `bson:"username" json:"username"`
	Email          string               `bson:"email" json:"email"`
	Password       string               `bson:"password" json:"password"`
	Bio            string               `bson:"bio" json:"bio"`
	FriendRequests []primitive.ObjectID `bson:"friendRequests" json:"friendRequests"`
	Friends        []primitive.ObjectID `bson:"friends" json:"friends"`
	CreatedAt      primitive.DateTime   `bson:"createdAt" json:"createdAt"`
}
