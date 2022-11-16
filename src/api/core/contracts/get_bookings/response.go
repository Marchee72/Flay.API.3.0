package get_bookings

import (
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/contracts/common"
	"flay-api-v3.0/src/api/core/entities"
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Bookings []booking `json:"bookings"`
}

type booking struct {
	ID          primitive.ObjectID `json:"id"`
	User        common.UserLw      `json:"user"`
	Apartment   common.Apartment   `json:"apartment"`
	CommonSpace string             `json:"common_space"`
	Building    common.BuildingLw  `json:"building"`
	Date        time.Time          `json:"date"`
	Shift       constants.Shift    `json:"shift"`
}

func (resp *Response) AddBooking(b entities.Booking, a lw.ApartmentLw) {
	newBooking := booking{
		ID:          b.Building.ID,
		User:        common.NewUserLw(b.User),
		Apartment:   common.NewApartment(a),
		CommonSpace: b.CommonSpace,
		Building:    common.NewBuildingLw(b.Building),
		Date:        b.Date,
		Shift:       b.Shift,
	}
	resp.Bookings = append(resp.Bookings, newBooking)
}
