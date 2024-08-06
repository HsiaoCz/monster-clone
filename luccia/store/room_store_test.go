package store

import (
	"context"
	"testing"

	"github.com/HsiaoCz/monster-clone/luccia/st"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateRoom(t *testing.T) {
	if err := Init(); err != nil {
		t.Fatal(err)
	}
	roomColl := client.Database(test_dbname).Collection("rooms")
	RoomStore := RoomStoreInit(client, roomColl)
	create_room_params := st.CreateRoomParams{
		Size:      "small",
		Seaside:   false,
		HotelID:   primitive.NewObjectID().String(),
		Available: true,
		BasePrice: 99.99,
		Price:     90.99,
	}
	room, err := st.NewRoomFromParams(create_room_params)
	if err != nil {
		t.Fatal(err)
	}
	roomResult, err := RoomStore.CreateRoom(context.Background(), room)
	if err != nil {
		t.Fatal(err)
	}
	if roomResult.Size != room.Size {
		t.Error("damn it")
	}
}
