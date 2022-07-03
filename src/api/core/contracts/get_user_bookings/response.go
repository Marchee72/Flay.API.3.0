package get_user_bookings

import (
	"time"

	"flay-api-v3.0/src/api/core/contracts/common"
	"flay-api-v3.0/src/api/core/entities"
)

type Response struct {
	Bookings []booking `json:"bookings"`
}

type booking struct {
	User        common.UserLw     `json:"user"`
	CommonSpace string            `json:"common_space"`
	Building    common.BuildingLw `json:"building"`
	StartDate   time.Time         `json:"start_date"`
	EndDate     time.Time         `json:"end_date"`
}

func (resp *Response) AddBooking(b entities.Booking) {
	newBooking := booking{
		User:        common.NewUserLw(b.User),
		CommonSpace: b.CommonSpace,
		Building:    common.NewBuildingLw(b.Building),
		StartDate:   b.StartDate,
		EndDate:     b.EndDate,
	}
	resp.Bookings = append(resp.Bookings, newBooking)
}
