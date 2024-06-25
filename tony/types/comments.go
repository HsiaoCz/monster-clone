package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comments struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"userID" json:"userID"`
	PostID     primitive.ObjectID `bson:"postID" json:"postID"`
	ParentID   primitive.ObjectID `bson:"parentID,omitempty" json:"parentID,omitempty"`
	Content    string             `bson:"content" json:"content"`
	CreateTime string             `bson:"createTime" json:"createTime"`
	Location   string             `bson:"location" json:"location"`
	Likes      string             `bson:"likes" json:"likes"`
	Unlikes    string             `bson:"unlikes" json:"unlikes"`
}

// create comments params
// need user login
type CreateCommentsParams struct {
	UserID     string `json:"userID"`
	PostID     string `json:"postID"`
	ParentID   string `json:"parentID,omitempty"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
	Location   string `json:"localtion"`
}

func (param CreateCommentsParams) Validate() map[string]string {
	errors := map[string]string{}
	if err := isValidID(param.UserID); err != nil {
		errors["userID"] = "invalid userID"
	}
	if err := isValidID(param.PostID); err != nil {
		errors["postID"] = "invalid postID"
	}
	if param.ParentID != "" {
		if err := isValidID(param.ParentID); err != nil {
			errors["parentID"] = "invalid parentID"
		}
	}
	return errors
}

func isValidID(id string) error {
	_, err := primitive.ObjectIDFromHex(id)
	return err
}

func NewCommentFromParams(params CreateCommentsParams) *Comments {
	comments := &Comments{}
	userID, _ := primitive.ObjectIDFromHex(params.UserID)
	postID, _ := primitive.ObjectIDFromHex(params.PostID)
	comments = &Comments{
		UserID:     userID,
		PostID:     postID,
		Content:    params.Content,
		CreateTime: params.CreateTime,
		Location:   params.Location,
		Likes:      "0",
		Unlikes:    "0",
	}
	if params.ParentID != "" {
		comments.ParentID, _ = primitive.ObjectIDFromHex(params.ParentID)
	}
	return comments
}
