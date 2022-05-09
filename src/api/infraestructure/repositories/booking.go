package store

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingRepository struct {
	Bookings *mongo.Collection
}

func (repository *BookingRepository) Save(ctx context.Context, booking entities.Booking) error {
	var err error
	_, err = repository.Bookings.InsertOne(ctx, booking)
	return err
}

func (repository *BookingRepository) IsAbailable(ctx context.Context) (bool, error) {
	var result entities.Booking
	err := repository.Bookings.FindOne(ctx, bson.M{"building.id": nil, "start_date": nil, "end_date": nil}).Decode(&result)
}
