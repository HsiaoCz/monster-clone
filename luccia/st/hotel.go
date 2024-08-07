package st

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string               `bson:"name" json:"name"`
	Localtion string               `bson:"location" json:"location"`
	Rooms     []primitive.ObjectID `bson:"rooms" json:"rooms"`
	Rating    int                  `bson:"rating" json:"rating"`
}

// type RoomType int

// const (
// 	_ RoomType = iota
// 	SingleRoomType
// 	DoubleRoomType
// 	SeaSideRoomType
// 	DeluxeRoomType
// )

type Room struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Size      string             `bson:"size" json:"size"`
	Seaside   bool               `bson:"seasize" json:"seasize"`
	BasePrice float64            `bson:"basePrice" json:"basePrice"`
	Price     float64            `bson:"price" json:"price"`
	HotelID   primitive.ObjectID `bson:"hotelID" json:"hotelID"`
	Available bool               `bson:"available" json:"available"`
}

type CreateRoomParams struct {
	Size      string  `json:"size"`
	Seaside   bool    `json:"seaside"`
	BasePrice float64 `json:"basePrice"`
	Price     float64 `json:"price"`
	HotelID   string  `json:"hotelOD"`
	Available bool    `json:"available"`
}

func NewRoomFromParams(parmas CreateRoomParams) (*Room, error) {
	HotelID, err := primitive.ObjectIDFromHex(parmas.HotelID)
	if err != nil {
		return nil, err
	}
	return &Room{
		Size:      parmas.Size,
		Seaside:   parmas.Seaside,
		Available: parmas.Available,
		BasePrice: parmas.BasePrice,
		Price:     parmas.Price,
		HotelID:   HotelID,
	}, nil
}

type UpdateHotelParams struct{}
type UpdateRoomParams struct {
	BasePrice float64 `bson:"basePrice" json:"basePrice"`
	Price     float64 `bson:"price" json:"price"`
	Available bool    `bson:"available" json:"available"`
}

type CreateHotelParmas struct {
	Name      string               `json:"name"`
	Localtion string               `json:"localtion"`
	Rooms     []primitive.ObjectID `json:"rooms"`
}
