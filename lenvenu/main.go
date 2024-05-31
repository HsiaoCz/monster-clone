package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/monster-clone/lenvenu/db"
	"github.com/HsiaoCz/monster-clone/lenvenu/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	ctx, cancl := context.WithTimeout(context.Background(), time.Second*5)
	defer cancl()
	clinet, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		log.Fatal(err)
	}
	var (
		logger         = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
		userColl       = clinet.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLL"))
		mongoUserStore = db.NewMongoUserStore(clinet, userColl)
		store          = &db.Store{User: mongoUserStore}
		userHandlers   = handlers.NewUserHandlers(store)
		app            = chi.NewRouter()
		srv            = http.Server{
			Handler:      app,
			Addr:         os.Getenv("PORT"),
			ReadTimeout:  time.Millisecond * 1500,
			WriteTimeout: time.Millisecond * 1500,
		}
	)
	slog.SetDefault(logger)

	app.Route("/app/v1", func(r chi.Router) {
		r.Post("/user", handlers.TransferHandlerFunc(userHandlers.HandleCreateUser))
	})

	slog.Info("the http server is running", "address", os.Getenv("PORT"))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	slog.Info("http server shutdown")
}
