package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSchema struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Username string `json:"username"`
	PublicKey string `json:"publicKey"`
}