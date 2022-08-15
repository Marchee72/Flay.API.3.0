package get_building_bookings

import (
	"context"
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/contracts/get_building_bookings"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request get_building_bookings.Request) (*get_building_bookings.Response, error)
}

type Implementation struct {
	BookingRepository providers.BookingRepository
}

func (usecase Implementation) Execute(ctx context.Context, request get_building_bookings.Request) (*get_building_bookings.Response, error) {
	from := time.Now().Add(time.Duration(constants.BookingMaxDaysAgo))
	bookings, err := usecase.BookingRepository.GetBuildingBookings(ctx, request.ID(), &from, request.Date)
	if err != nil {
		return nil, err
	}
	response := get_building_bookings.Response{}
	for _, b := range bookings {
		response.AddBooking(b)
	}
	return &response, nil
}
