package api

import (
	"context"
	"github.com/sergio9875/hotel-service/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

const (
	testDbUri  = "mongodb://localhost:27017"
	testDbName = "hotel-reservation-test"
)

type testDb struct {
	client *mongo.Client
	*db.Store
}

func (tdb *testDb) teardown(t *testing.T) {
	if err := tdb.client.Database(db.DbName).Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testDb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDbUri))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client)
	return &testDb{
		client: client,
		Store: &db.Store{
			Hotel:   hotelStore,
			User:    db.NewMongoUserStore(client),
			Room:    db.NewMongoRoomStore(client, hotelStore),
			Booking: db.NewMongoBookingStore(client),
		},
	}
}
