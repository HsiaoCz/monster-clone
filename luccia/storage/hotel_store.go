package storage

import (
	"context"
	"errors"

	"github.com/HsiaoCz/monster-clone/luccia/st"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStorer interface {
	CreateHotel(context.Context, *st.Hotel) (*st.Hotel, error)
	DeleteHotel(context.Context, primitive.ObjectID) error
	GetHotels(context.Context, bson.M) ([]*st.Hotel, error)
	GetHotelByID(context.Context, primitive.ObjectID) (*st.Hotel, error)
	UpdateHotel(context.Context, bson.M, bson.M) (*st.Hotel, error)
}

type HotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func HotelStoreInit(client *mongo.Client, coll *mongo.Collection) *HotelStore {
	return &HotelStore{
		client: client,
		coll:   coll,
	}
}

func (h *HotelStore) CreateHotel(ctx context.Context, hotel *st.Hotel) (*st.Hotel, error) {
	filter := bson.D{
		{Key: "$get", Value: bson.D{
			{Key: "name", Value: hotel.Name},
			{Key: "localtion", Value: hotel.Localtion},
		}},
	}
	result := h.coll.FindOne(ctx, filter)
	if result.Err() != mongo.ErrNoDocuments {
		return nil, errors.New("this record exists")
	}
	res, err := h.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (h *HotelStore) DeleteHotel(ctx context.Context, hotelID primitive.ObjectID) error {
	res, err := h.coll.DeleteOne(ctx, bson.M{"_id": hotelID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no this record")
	}
	return nil
}

func (h *HotelStore) GetHotels(ctx context.Context, filter bson.M) ([]*st.Hotel, error) {
	var hotels []*st.Hotel
	res, err := h.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := res.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}

func (h *HotelStore) GetHotelByID(ctx context.Context, hotelID primitive.ObjectID) (*st.Hotel, error) {
	return nil, nil
}

func (h *HotelStore) UpdateHotel(ctx context.Context, filter bson.M, update bson.M) (*st.Hotel, error) {
	return nil, nil
}
