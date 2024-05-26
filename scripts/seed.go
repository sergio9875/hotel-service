package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sergio9875/hotel-service/db"
	"github.com/sergio9875/hotel-service/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DbUri))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DbName)
	roomStore := db.NewMongoRoomStore(client, db.DbName)

	hotel := types.Hotel{
		Name:     "Lucia",
		Location: "Spain",
	}
	rooms := []types.Room{

		{
			Type:      types.SingleRoomType,
			BasePrice: 99.9,
		},
		{
			Type:      types.DeluxeRoomType,
			BasePrice: 199.9,
		},
		{
			Type:      types.SeaSideRoomType,
			BasePrice: 122.9,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted-room", insertedRoom)
	}
}
