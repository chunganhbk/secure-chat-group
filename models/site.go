package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SiteSchema struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name string `json:"name"`
	Secret string `json:"secret"`
}






