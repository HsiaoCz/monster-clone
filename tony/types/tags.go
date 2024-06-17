package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Content string             `bson:"content" json:"content"`
}

type CreateTagParams struct {
	Content string `json:"content"`
}

var (
	minContentLen = 6
	maxContentLen = 12
)

func (param CreateTagParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(param.Content) < minContentLen || len(param.Content) > maxContentLen {
		errors["content"] = fmt.Sprintf("content shouldn't less then %d or more then %d", minContentLen, maxContentLen)
	}
	return errors
}

func TagFromParams(param CreateTagParams) *Tag {
	return &Tag{
		Content: param.Content,
	}
}
