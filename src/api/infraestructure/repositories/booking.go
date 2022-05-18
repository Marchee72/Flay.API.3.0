package store

import (
	"context"
	"time"

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
	var response []entities.Booking
	if err = cursor.All(ctx, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (repository *BookingRepository) IsAbailable(ctx context.Context, buildingID primitive.ObjectID, startDate time.Time, endDate time.Time) (bool, error) {
	var result entities.Booking
	filter := bson.M{
		"building.id": buildingID,
		"$or": bson.A{
			bson.A{"$and",
				bson.M{"start_date": bson.M{"$gte": startDate}},
				bson.M{"start_date": bson.M{"$lte": endDate}},
			},
			bson.A{"$and",
				bson.M{"end_date": bson.M{"$gte": startDate}},
				bson.M{"end_date": bson.M{"$lte": endDate}},
			},
			bson.A{"$and",
				bson.M{"start_date": bson.M{"$lte": startDate}},
				bson.M{"end_date": bson.M{"$gte": startDate}},
			},
			bson.A{"$and",
				bson.M{"start_date": bson.M{"$lte": endDate}},
				bson.M{"end_date": bson.M{"$gte": endDate}},
			},
		},
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
