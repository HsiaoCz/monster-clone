package store

import (
	"context"

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
	return nil, nil
}

func (r *RoomStore) DeleteRoom(ctx context.Context, roomID primitive.ObjectID) error {
	return nil
}

func (r *RoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*st.Room, error) {
	return nil, nil
}

func (r *RoomStore) GetRoomByID(ctx context.Context, roomID primitive.ObjectID) (*st.Room, error) {
	return nil, nil
}

func (r *RoomStore) UpdateRoom(ctx context.Context, roomID primitive.ObjectID, updateRoomParams *st.UpdateRoomParams) (*st.Room, error) {
	return nil, nil
}
