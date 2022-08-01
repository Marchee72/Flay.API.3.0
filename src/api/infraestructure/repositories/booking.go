package store

import (
	"context"
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repository *BookingRepository) GetUserBookings(ctx context.Context, userID primitive.ObjectID) ([]entities.Booking, error) {
	cursor, err := repository.Bookings.Find(ctx, bson.M{"user._id": userID})
	if err != nil {
		return nil, err
	}
	var result []entities.Booking
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (repository *BookingRepository) GetBuildingBookings(ctx context.Context, buildingID primitive.ObjectID, from time.Time) ([]entities.Booking, error) {
	fromBson := primitive.NewDateTimeFromTime(from)
	cursor, err := repository.Bookings.Find(ctx, bson.M{"building._id": buildingID, "date": bson.M{"$gte": fromBson}})
	if err != nil {
		return nil, err
	}
	var result []entities.Booking
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (repository *BookingRepository) IsAbailable(ctx context.Context, buildingID primitive.ObjectID, commonSpace string, date time.Time, shift constants.Shift) (bool, error) {
	var result entities.Booking
	filter := bson.M{
		"building.id":  buildingID,
		"common_space": commonSpace,
		"date":         date,
		"shift":        shift,
	}
	err := repository.Bookings.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return true, nil
		}
		return false, err
	}
	return true, nil
}
