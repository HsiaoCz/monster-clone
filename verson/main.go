package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/monster-clone/verson/handlers"
	"github.com/HsiaoCz/monster-clone/verson/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// file,err:=os.OpenFile("verson.log",)
	if err := logger.InitLogger(); err != nil {
		log.Fatal(err)
	}

	var (
		port         = os.Getenv("PORT")
		userHandlers = &handlers.UserHandlers{}
		app          = http.NewServeMux()
	)
	{
		// router
		app.HandleFunc("GET /hello", handlers.TransferHandlerFunc(userHandlers.HandleCreateUser))
	}
	logger.Logger.Info("the http server is running", "listen address", port)
	http.ListenAndServe(port, app)
}
