package app

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestCreatePost(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
}

func TestDeletePosByID(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
}
