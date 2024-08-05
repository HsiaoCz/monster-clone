package store

import (
	"context"

	"github.com/HsiaoCz/monster-clone/luccia/st"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingStorer interface {
	GetBookings(context.Context, primitive.ObjectID) ([]*st.Booking, error)
	UpdateBooking(context.Context, primitive.ObjectID, *st.UpdateBookingParams) (*st.Booking, error)
	CreateBooking(context.Context, *st.Booking) (*st.Booking, error)
	GetBookingByID(context.Context, primitive.ObjectID) (*st.Booking, error)
}

type BookingStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func BookingStoreInit(client *mongo.Client, coll *mongo.Collection) *BookingStore {
	return &BookingStore{
		client: client,
		coll:   coll,
	}
}

func (b *BookingStore) GetBookings(context.Context, primitive.ObjectID) ([]*st.Booking, error) {
	return nil, nil
}
func (b *BookingStore) UpdateBooking(context.Context, primitive.ObjectID, *st.UpdateBookingParams) (*st.Booking, error) {
	return nil, nil
}
func (b *BookingStore) CreateBooking(context.Context, *st.Booking) (*st.Booking, error) {
	return nil, nil
}
func (b *BookingStore) GetBookingByID(context.Context, primitive.ObjectID) (*st.Booking, error) {
	return nil, nil
}
