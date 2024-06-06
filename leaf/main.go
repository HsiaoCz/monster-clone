package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	v1 "github.com/HsiaoCz/monster-clone/leaf/app/v1"
	"github.com/HsiaoCz/monster-clone/leaf/store"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal(err)
	// }
	// env is totally shit
	// var (
	// 	user     = os.Getenv("MUSER")
	// 	password = os.Getenv("PASSWORD")
	// 	host     = os.Getenv("HOST")
	// 	port     = os.Getenv("PORT")
	// 	dbname   = os.Getenv("DBNAME")
	// )
	// we don't need this shit
	// but if the db connect error we use this to check shit where?
	// fmt.Println(user, password, host, port, dbname)
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI())
	if err != nil {
		log.Fatal(err)
	}
 
	

	var (
		logger         = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
		mysqlUserStore = store.NewMongoUserStore()
		store          = &store.Store{User: mysqlUserStore}
		userHandlers   = v1.NewUserAPI(store)
		router         = fiber.New()
		av1            = router.Group("/app/v1")
	)
	slog.SetDefault(logger)
	// routers
	{
		av1.Post("/user")
	}
	router.Listen(":3001")
}
