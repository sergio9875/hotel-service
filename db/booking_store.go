package db

import (
	"context"
	"github.com/sergio9875/hotel-service/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingStore interface {
	InsertBooking(ctx context.Context, booking *types.Booking) (*types.Booking, error)
	GetBookings(ctx context.Context, m bson.M) ([]*types.Booking, error)
}

type MongoBookingStore struct {
	client *mongo.Client
	coll   *mongo.Collection

	BookingStore
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	return &MongoBookingStore{
		client: client,
		coll:   client.Database(DbName).Collection("bookings"),
	}
}
func (s *MongoBookingStore) GetBookings(ctx context.Context, filter bson.M) ([]*types.Booking, error) {
	curr, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var bookings []*types.Booking
	if err := curr.All(ctx, &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (s *MongoBookingStore) InsertBooking(ctx context.Context, booking *types.Booking) (*types.Booking, error) {
	resp, err := s.coll.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}
	booking.ID = resp.InsertedID.(primitive.ObjectID)
	return booking, nil
}
