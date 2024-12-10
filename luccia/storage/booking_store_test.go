package storage

import "testing"

func TestBookingRoom(t *testing.T) {
	if err := Init(); err != nil {
		t.Fatal(err)
	}
}
