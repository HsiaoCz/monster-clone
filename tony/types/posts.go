package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Posts struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	UserID      primitive.ObjectID `bson:"userID" json:"userID"`
	Title       string             `bson:"title" json:"title"`
	Content     string             `bson:"content" json:"content"`
	CreateTime  time.Time          `bson:"createTime" json:"createTime"`
	Since       string             `bson:"since" json:"since"`
	Location    string             `bson:"localtion" json:"localtion"`
	Likes       string             `bson:"likes" json:"likes"`
	Collections string             `bson:"collections" json:"collections"`
	Tags        []string           `bson:"tags" json:"tags"`
	Classify    []string           `bson:"classify" json:"classify"`
	Comments    string             `bson:"comments" json:"comments"`
}

type CreatePostsParams struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Tags     []string `json:"tags"`
	Classify []string `json:"classify"`
}

var (
	minTitleLen = 6
	maxTitleLen = 200
)

func (params CreatePostsParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.Title) < minTitleLen || len(params.Title) > maxTitleLen {
		errors["title"] = fmt.Sprintf("the title len should less than %d and more then %d", maxTitleLen, minTitleLen)
	}
	return errors
}
