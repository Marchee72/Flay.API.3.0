package get_user_bookings

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/get_bookings"
	"flay-api-v3.0/src/api/core/contracts/get_user_bookings"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request get_user_bookings.Request) (*get_bookings.Response, error)
}

type Implementation struct {
	BookingRepository providers.BookingRepository
	UserRepository    providers.UserRepository
}

func (usecase Implementation) Execute(ctx context.Context, request get_user_bookings.Request) (*get_bookings.Response, error) {
	bookings, err := usecase.BookingRepository.GetUserBookings(ctx, request.UserID())
	if err != nil {
		return nil, err
	}
	user, err := usecase.UserRepository.GetUserById(ctx, request.UserID())
	if err != nil {
		return nil, err
	}
	response := get_bookings.Response{}
	for _, b := range bookings {
		response.AddBooking(b, *user.Apartment)
	}
	return &response, nil
}
