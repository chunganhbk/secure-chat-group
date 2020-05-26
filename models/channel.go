package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChannelSchema struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name string
	PrivateKeys map[string]string // [userID]: Channels symmetric AES key is encrypted with the select users public key
	UserMap map[string]string
}