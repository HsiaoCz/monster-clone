package store

import (
	"context"
	"errors"

	"github.com/HsiaoCz/monster-clone/luccia/st"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStorer interface {
	CreateRoom(context.Context, *st.Room) (*st.Room, error)
	DeleteRoom(context.Context, primitive.ObjectID) error
	GetRooms(context.Context, bson.M) ([]*st.Room, error)
	GetRoomByID(context.Context, primitive.ObjectID) (*st.Room, error)
	UpdateRoom(context.Context, primitive.ObjectID, *st.UpdateRoomParams) (*st.Room, error)
}

type RoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func RoomStoreInit(client *mongo.Client, coll *mongo.Collection) *RoomStore {
	return &RoomStore{
		client: client,
		coll:   coll,
	}
}

func (r *RoomStore) CreateRoom(ctx context.Context, room *st.Room) (*st.Room, error) {
	result, err := r.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = result.InsertedID.(primitive.ObjectID)
	return room, nil
}

func (r *RoomStore) DeleteRoom(ctx context.Context, roomID primitive.ObjectID) error {
	result, err := r.coll.DeleteOne(ctx, bson.M{"_id": roomID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("database don't have this record")
	}
	return nil
}

func (r *RoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*st.Room, error) {
	var rooms []*st.Room
	result, err := r.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := result.Decode(&rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *RoomStore) GetRoomByID(ctx context.Context, roomID primitive.ObjectID) (*st.Room, error) {
	var room st.Room
	if err := r.coll.FindOne(ctx, bson.M{"_id": roomID}).Decode(&room); err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomStore) UpdateRoom(ctx context.Context, roomID primitive.ObjectID, updateRoomParams *st.UpdateRoomParams) (*st.Room, error) {
	return nil, nil
}
