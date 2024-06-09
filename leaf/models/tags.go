package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tags struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	Content string             `bson:"content" json:"content"`
}
