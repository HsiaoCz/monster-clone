package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/monster-clone/fetch/handlers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	var (
		port         = os.Getenv("PORT")
		priceHandler = &handlers.PriceHandler{}
		router       = http.NewServeMux()
	)

	router.HandleFunc("GET /fetch", handlers.TransferHandlerfunc(priceHandler.HandleFetchPrice))

	logrus.WithFields(logrus.Fields{
		"listen address": port,
	}).Info("http server is running")

	http.ListenAndServe(port, router)
}
