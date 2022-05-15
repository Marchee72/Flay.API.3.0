package get_user_bookings

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/get_user_bookings"
	"flay-api-v3.0/src/api/core/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Execute(ctx context.Context, userID primitive.ObjectID) (*get_user_bookings.Response, error)
}

type Implementation struct {
	BookingRepository providers.BookingRepository
}

func (usecase Implementation) Execute(ctx context.Context, userID primitive.ObjectID) (*get_user_bookings.Response, error) {
	bookings, err := usecase.BookingRepository.GetUserBookings(ctx, userID)
	if err != nil {
		return nil, err
	}
	response := get_user_bookings.Response{}
	for _, b := range bookings {
		response.AddBooking(b)
	}
	return &response, nil
}
