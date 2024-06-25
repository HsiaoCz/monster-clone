package store

import (
	"context"
	"os"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentStorer interface {
	CreateComment(context.Context, *types.Comments) (*types.Comments, error)
	DeleteCommentByID(context.Context, primitive.ObjectID) error
	GetCommentsByPostID(context.Context, primitive.ObjectID) ([]*types.Comments, error)
	// get comments by postID and parentID
	GetCommentsByPostIDAndParentID(context.Context, primitive.ObjectID, primitive.ObjectID) ([]*types.Comments, error)
}

type MongoCommentStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoCommentStore(client *mongo.Client, coll *mongo.Collection) *MongoCommentStore {
	return &MongoCommentStore{
		client: client,
		coll:   client.Database(os.Getenv("DBNAME")).Collection("comments"),
	}
}

func (m *MongoCommentStore) CreateComment(ctx context.Context, comment *types.Comments) (*types.Comments, error) {
	// create comments dont need to check exist
	result, err := m.coll.InsertOne(ctx, comment)
	if err != nil {
		return nil, err
	}
	comment.ID = result.InsertedID.(primitive.ObjectID)
	return comment, nil
}
func (m *MongoCommentStore) DeleteCommentByID(ctx context.Context, cid primitive.ObjectID) error {
	filter := bson.D{
		{Key: "_id", Value: cid},
	}
	_, err := m.coll.DeleteOne(ctx, filter)
	return err
}

func (m *MongoCommentStore) GetCommentsByPostID(ctx context.Context, pid primitive.ObjectID) ([]*types.Comments, error) {
	filter := bson.D{
		{Key: "postID", Value: pid},
	}
	var comments []*types.Comments
	cusor, err := m.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cusor.All(ctx, comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func (m *MongoCommentStore) GetCommentsByPostIDAndParentID(ctx context.Context, postID primitive.ObjectID, parentID primitive.ObjectID) ([]*types.Comments, error) {
	filter := bson.M{
		"postID":   postID,
		"parentID": parentID,
	}
	var comments []*types.Comments
	cusor, err := m.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cusor.All(ctx, comments); err != nil {
		return nil, err
	}
	return comments, nil
}
