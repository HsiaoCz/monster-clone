package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comments struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	UserID     primitive.ObjectID `bson:"userID" json:"userID"`
	PostID     primitive.ObjectID `bson:"postID" json:"postID"`
	ParentID   primitive.ObjectID `bson:"parentID" json:"parentID"`
	Content    string             `bson:"content" json:"content"`
	CreateTime string             `bson:"createTime" json:"createTime"`
	Location   string             `bson:"location" json:"location"`
	Likes      string             `bson:"likes" json:"likes"`
}

// create comments params
// need user login
type CreateCommentsParams struct {
	UserID     primitive.ObjectID `json:"userID"`
	PostID     primitive.ObjectID `json:"postID"`
	ParentID   primitive.ObjectID `json:"parentID"`
	Content    string             `json:"content"`
	CreateTime string             `json:"createTime"`
	Likes      string             `json:"likes"`
}
