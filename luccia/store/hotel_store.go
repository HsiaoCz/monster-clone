package store

import (
	"context"

	"github.com/HsiaoCz/monster-clone/luccia/st"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelStorer interface {
	CreateHotel(context.Context, *st.Hotel) (*st.Hotel, error)
	DeleteHotel(context.Context, primitive.ObjectID) error
	GetHotels(context.Context, bson.M) ([]*st.Hotel, error)
	GetHotelByID(context.Context, primitive.ObjectID) (*st.Hotel, error)
	UpdateHotel(context.Context, primitive.ObjectID) (*st.Hotel, error)
}
