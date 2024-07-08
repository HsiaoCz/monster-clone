package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/monster-clone/verson/handlers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// file,err:=os.OpenFile("verson.log",)
	var (
		port         = os.Getenv("PORT")
		userHandlers = &handlers.UserHandlers{}
		app          = http.NewServeMux()
	)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	{
		// router
		app.HandleFunc("GET /hello", handlers.TransferHandlerFunc(userHandlers.HandleCreateUser))
	}

	logrus.WithFields(logrus.Fields{
		"listen address": port,
	}).Info("http server is running")
	
	http.ListenAndServe(port, app)
}
