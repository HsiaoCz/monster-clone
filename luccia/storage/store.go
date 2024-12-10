package storage

type Store struct {
	Us    UserStorer
	Room  RoomStorer
	Hotel HotelStorer
	Book  BookingStorer
}
