package app

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestAPPHello(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
	t.Log(os.Getenv("PORT"))
}
