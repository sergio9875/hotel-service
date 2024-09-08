package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/sergio9875/hotel-service/db"
	"github.com/sergio9875/hotel-service/db/fixtures"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var (
		ctx           = context.Background()
		mongoEndpoint = os.Getenv("MONGO_DB_URL")
		mongoDBName   = os.Getenv("MONGO_DB_NAME")
	)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(mongoDBName).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client)

	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Booking: db.NewMongoBookingStore(client),
		Room:    db.NewMongoRoomStore(client, hotelStore),
		Hotel:   hotelStore,
	}

	user := fixtures.AddUser(store, "Sergey", "Jameson", false)
	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println("user ->", user)
	hotel := fixtures.AddHotel(store, "Lucia", "Spain", 3, nil)
	hotel2 := fixtures.AddHotel(store, "Paradise", "France", 4, nil)
	room := fixtures.AddRoom(store, "large", true, 88.64, hotel.ID)
	room2 := fixtures.AddRoom(store, "king-size", true, 127.55, hotel2.ID)
	fmt.Println("admin ->", admin)
	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 5))
	fixtures.AddBooking(store, user.ID, room2.ID, time.Now(), time.Now().AddDate(0, 0, 3))
	fmt.Println("booking ->", booking)

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("random hotel name %d", i)
		location := fmt.Sprintf("location %d", i)
		fixtures.AddHotel(store, name, location, rand.Intn(5)+1, nil)
	}
}
