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
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestCreatePost(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
	if err := db.Init(); err != nil {
		t.Fatal(err)
	}

	postHandler := NewPostHandler(data.NewPostStore(db.Get()))

	app := fiber.New()

	app.Post("/post", postHandler.HandleCreatePost)

	parmas := types.CreatePostParams{
		UserID:   uuid.New().String(),
		Content:  "something",
		PostPath: "./post/1234.txt",
	}

	b, _ := json.Marshal(parmas)
	req := httptest.NewRequest("POST", "/post", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	postParams := types.CreatePostParams{}

	json.NewDecoder(resp.Body).Decode(&postParams)

	if parmas.Content != postParams.Content {
		t.Errorf("want %s but got %s", parmas.Content, postParams.Content)
	}

}
