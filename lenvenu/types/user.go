package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Username string             `bson:"username" json:"username"`
}