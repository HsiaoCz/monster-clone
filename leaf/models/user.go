package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username   string             `bson:"username" json:"username"`
	Password   string             `bson:"password" json:"-"`
	Level      int                `bson:"level" json:"level"`
	Content    string             `bson:"content" json:"content"`
	Email      string             `bson:"email" json:"email"`
	Phone      string             `bson:"phone" json:"phone"`
	Job        string             `bson:"job" json:"job"`
	Company    string             `bson:"company" json:"company"`
	Birthday   string             `bson:"birthday" json:"birthday"`
	Age        int                `bson:"age" json:"age"`
	Avatar     string             `bson:"avatar" json:"avatar"`
	Tags       []int64            `bson:"tags" json:"tags"`
	Describe   int                `bson:"describe" json:"describe"`
	Collection int                `bson:"collection" json:"collection"`
}
