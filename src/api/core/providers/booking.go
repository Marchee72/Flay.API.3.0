package providers

import (
	"context"
	"time"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingRepository interface {
	Save(ctx context.Context, booking entities.Booking) error
	IsAbailable(ctx context.Context, buildingID primitive.ObjectID, startDate time.Time, endDate time.Time) (bool, error)
}
