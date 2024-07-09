package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Posts struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	UserID      primitive.ObjectID `bson:"userID" json:"userID"`
	Title       string             `bson:"title" json:"title"`
	Content     string             `bson:"content" json:"content"`
	CreateTime  string             `bson:"createTime" json:"createTime"`
	Since       string             `bson:"since" json:"since"`
	Location    string             `bson:"localtion" json:"localtion"`
	Likes       string             `bson:"likes" json:"likes"`
	Collections string             `bson:"collections" json:"collections"`
	Tags        []string           `bson:"tags" json:"tags"`
	Classify    []string           `bson:"classify" json:"classify"`
}
