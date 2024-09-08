package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sergio9875/hotel-service/api"
	"github.com/sergio9875/hotel-service/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var config = fiber.Config{
	ErrorHandler: api.ErrorHandler,
}

func main() {
	mongoEndpoint := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	//MySql client
	//_, err = db.ConnectToMySql(db.MySqlUser, db.MySqlPassword, db.MySqlHost, db.DbName)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("MySql connected..")

	//Postgres client
	//postgresDb, err := sql.Open("postgres", db.PostgresConnString)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Handlers initialization
	var (
		hotelsStore  = db.NewMongoHotelStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelsStore)
		userStore    = db.NewMongoUserStore(client)
		bookingStore = db.NewMongoBookingStore(client)
		store        = &db.Store{
			Hotel:   hotelsStore,
			User:    userStore,
			Room:    roomStore,
			Booking: bookingStore,
		}
		userHandler    = api.NewUserHandler(userStore)
		hotelHandler   = api.NewHotelHandler(store)
		authHandler    = api.NewAuthHandler(userStore)
		roomHandler    = api.NewRoomHandler(store)
		bookingHandler = api.NewBookingHandler(store)
		app            = fiber.New(config)
		auth           = app.Group("/api")
		apiV1          = app.Group("/api/v1", api.JWTAuthentication(userStore))
		admin          = apiV1.Group("/admin", api.AdminAuth)
	)
	// Auth
	auth.Post("/auth", authHandler.HandleAuthenticate)

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

	// Rooms handlers
	apiV1.Post("/room/:id/book", roomHandler.HandleBookRoom)
	apiV1.Get("/room/", roomHandler.HandleGetRooms)

	// Booking handlers
	apiV1.Get("/booking/:id", bookingHandler.HandleGetBooking)
	apiV1.Get("/booking/:id/cancel", bookingHandler.HandleCancelBooking)

	// Admin handlers
	admin.Get("/booking/", bookingHandler.HandleGetBookings)

	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
	err = app.Listen(listenAddr)
	if err != nil {
		return
	}

}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
