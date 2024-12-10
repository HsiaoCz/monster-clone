package storage

import "testing"

func TestCreateHotel(t *testing.T) {
	if err := Init(); err != nil {
		t.Fatal(err)
	}
	hotelColl := client.Database(test_dbname).Collection("hotels")
	hotelStore := HotelStoreInit(client, hotelColl)

	_ = hotelStore
}
