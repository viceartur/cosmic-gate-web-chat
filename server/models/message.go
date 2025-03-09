package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	// SenderID    primitive.ObjectID `bson:"senderId" json:"senderId"`
	// RecipientID primitive.ObjectID `bson:"recipientId" json:"recipientId"`
	SenderID    int       `bson:"senderId" json:"senderId"`
	RecipientID int       `bson:"recipientId" json:"recipientId"`
	Text        string    `bson:"text" json:"text"`
	SentAt      time.Time `bson:"sentAt" json:"sentAt"`
}
