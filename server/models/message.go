package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SenderID    string             `bson:"senderId" json:"senderId"`
	RecipientID string             `bson:"recipientId" json:"recipientId"`
	Text        string             `bson:"text" json:"text"`
	SentAt      time.Time          `bson:"sentAt" json:"sentAt"`
}
