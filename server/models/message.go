package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SenderID    primitive.ObjectID `bson:"senderId" json:"senderId"`
	RecipientID primitive.ObjectID `bson:"recipientId" json:"recipientId"`
	Text        string             `bson:"text" json:"text"`
	SentAt      primitive.DateTime `bson:"sentAt" json:"sentAt"`
}
