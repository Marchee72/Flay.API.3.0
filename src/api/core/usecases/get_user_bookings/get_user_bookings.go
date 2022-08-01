package get_user_bookings

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/get_user_bookings"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request get_user_bookings.Request) (*get_user_bookings.Response, error)
}

type Implementation struct {
	BookingRepository providers.BookingRepository
}

func (usecase Implementation) Execute(ctx context.Context, request get_user_bookings.Request) (*get_user_bookings.Response, error) {
	bookings, err := usecase.BookingRepository.GetUserBookings(ctx, request.UserID)
	if err != nil {
		return nil, err
	}
	response := get_user_bookings.Response{}
	for _, b := range bookings {
		response.AddBooking(b)
	}
	return &response, nil
}
