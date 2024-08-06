package store

import "testing"

func TestCreateHotel(t *testing.T) {
	if err := Init(); err != nil {
		t.Fatal(err)
	}
}
