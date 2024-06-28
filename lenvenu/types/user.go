package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Username  string             `bson:"username" json:"username"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt string             `bson:"createdAt" json:"createdAt"`
}

// user context infomation
type UserInfo struct {
	UserID  primitive.ObjectID
	Email   string
	IsAdmin bool
}
