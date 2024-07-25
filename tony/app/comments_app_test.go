package app

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateComment(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteComment(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
}
