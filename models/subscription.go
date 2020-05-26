package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SubscriptionSchema struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID primitive.ObjectID
	Type string `json:"type"`
	Endpoint string `json:"end_point"`
	ExpirationTime string `json:"expiration_time"`
}
