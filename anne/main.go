package main

import (
	"log"
	"net/http"
	"os"

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
		port   = os.Getenv("PORT")
		router = http.NewServeMux()
	)

	router.HandleFunc("GET /price", func(w http.ResponseWriter, r *http.Request) {})
	logrus.WithFields(logrus.Fields{
		"port":   port,
		"auther": "HsiaoCz",
		"name":   "anne",
	}).Info("the http server is running")
	http.ListenAndServe(port, router)
}
