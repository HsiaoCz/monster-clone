package store

import (
	"context"

	"github.com/HsiaoCz/monster-clone/luccia/st"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingStorer interface {
	GetBookings(context.Context, primitive.ObjectID) ([]*st.Booking, error)
	UpdateBooking(context.Context, primitive.ObjectID, *st.UpdateBookingParams) (*st.Booking, error)
}
