package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HsiaoCz/monster-clone/santino/data"
	"github.com/HsiaoCz/monster-clone/santino/db"
	"github.com/HsiaoCz/monster-clone/santino/types"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func TestCreateComment(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
	if err := db.Init(); err != nil {
		t.Fatal(err)
	}

	userHandler := NewUserHandler(data.NewUserData(db.Get()))

	app := fiber.New()

	app.Post("/user", userHandler.HandleCreateUser)

	parmas := types.CreateUserParams{
		Username:         "shawcz",
		Email:            "shawcz@gmail.com",
		Password:         "shawcz123",
		Synopsis:         "something wrong",
		Avatar:           "./picture/1233.jpg",
		Background_Image: "./bgi/1234.jpg",
	}

	b, _ := json.Marshal(parmas)
	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	userParams := types.CreateUserParams{}

	json.NewDecoder(resp.Body).Decode(&userParams)
}
