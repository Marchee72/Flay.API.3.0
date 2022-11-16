package get_building_bookings

import (
	"context"
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/contracts/get_bookings"
	"flay-api-v3.0/src/api/core/contracts/get_building_bookings"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request get_building_bookings.Request) (*get_bookings.Response, error)
}

type Implementation struct {
	BookingRepository providers.BookingRepository
	UserRepository    providers.UserRepository
}

func (usecase Implementation) Execute(ctx context.Context, request get_building_bookings.Request) (*get_bookings.Response, error) {
	from := time.Now().Add(time.Duration(constants.BookingMaxDaysAgo))
	bookings, err := usecase.BookingRepository.GetBuildingBookings(ctx, request.BuildingID(), &from, request.Date)
	if err != nil {
		return nil, err
	}
	response := get_bookings.Response{}
	for _, b := range bookings {
		user, err := usecase.UserRepository.GetUserById(ctx, b.User.ID)
		if err != nil {
			return nil, err
		}
		response.AddBooking(b, *user.Apartment)
	}
	return &response, nil
}
