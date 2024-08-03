package store

import (
	"context"

	"github.com/HsiaoCz/monster-clone/luccia/st"
)

type BookingStorer interface{
	GetBookings(context.Context)([]*st.Booking,error)
}