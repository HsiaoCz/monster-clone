package store

import (
	"context"

	"github.com/HsiaoCz/monster-clone/luccia/st"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomStorer interface {
	CreateRoom(context.Context, *st.Room) (*st.Room, error)
	DeleteRoom(context.Context, primitive.ObjectID) error
	GetRooms(context.Context, string) ([]*st.Room, error)
	GetRoomByID(context.Context, primitive.ObjectID) (*st.Room, error)
	UpdateRoom(context.Context, primitive.ObjectID,*st.UpdateRoomParams) (*st.Room, error)
}
