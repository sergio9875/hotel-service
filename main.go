package main

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/sergio9875/hotel-service/api"
	"github.com/sergio9875/hotel-service/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":4000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DbUri))
	if err != nil {
		log.Fatal(err)
	}
	// Handlers initialization
	var (
		hotelsStore = db.NewMongoHotelStore(client)
		roomStore   = db.NewMongoRoomStore(client, hotelsStore)
		userStore   = db.NewMongoUserStore(client)
		store       = &db.Store{
			Hotel: hotelsStore,
			User:  userStore,
			Room:  roomStore,
		}
		userHandler  = api.NewUserHandler(userStore)
		hotelHandler = api.NewHotelHandler(store)
		app          = fiber.New(config)
		apiV1        = app.Group("/api/v1")
	)

	// User handlers
	apiV1.Put("user/:id", userHandler.HandlePutUser)
	apiV1.Delete("user/:id", userHandler.HandleDeleteUser)
	apiV1.Post("user", userHandler.HandlePostUser)
	apiV1.Get("/user", userHandler.HandleGetUsers)
	apiV1.Get("/user/:id", userHandler.HandleGetUser)

	// Hotel handlers
	apiV1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiV1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)
	apiV1.Get("/hotel/:id", hotelHandler.HandleGetHotel)

	err = app.Listen(*listenAddr)
	if err != nil {
		log.Fatal(err)
	}

}
