package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MessageSchema struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	ChannelID primitive.ObjectID `bson:"channel_id"`
	Timestamp time.Time
	Encrypted string
}
