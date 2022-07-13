package providers

import (
	"context"
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingRepository interface {
	Save(ctx context.Context, booking entities.Booking) error
	GetUserBookings(ctx context.Context, userID primitive.ObjectID) ([]entities.Booking, error)
	IsAbailable(ctx context.Context, buildingID primitive.ObjectID, commonSpace string, date time.Time, shift constants.Shift) (bool, error)
}
