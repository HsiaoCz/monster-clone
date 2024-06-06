package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Username   string               `bson:"username" json:"username"`
	Password   string               `bson:"password" json:"-"`
	Level      int                  `bson:"level" json:"level"`
	Content    string               `bson:"content" json:"content"`
	Email      string               `bson:"email" json:"email"`
	Phone      string               `bson:"phone" json:"phone"`
	Job        string               `bson:"job" json:"job"`
	Company    string               `bson:"company" json:"company"`
	Birthday   string               `bson:"birthday" json:"birthday"`
	Age        string               `bson:"age" json:"age"`
	Avatar     string               `bson:"avatar" json:"avatar"`
	Tags       []primitive.ObjectID `bson:"tags" json:"tags"`
	Describe   int                  `bson:"describe" json:"describe"`
	Collection int                  `bson:"collection" json:"collection"`
	IsAdmin    bool                 `bson:"isAdmin" json:"isAdmin"`
}

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone"`
	Birthday string `json:"birthday"`
}

type UserInfo struct {
	UserID  primitive.ObjectID
	Email   string
	IsAdmin bool
}
