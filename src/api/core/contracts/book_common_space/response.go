package book_common_space

import (
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/entities"
)

type Response struct {
	Booking            booking `json:"booking"`
	HasActivePenalties bool    `json:"has_active_penalties"`
	Penalty            penalty `json:"penalty"`
	IsAbailable        bool    `json:"is_abailable"`
}

func (resp *Response) SetPenalty(p entities.Penalty) {
	resp.HasActivePenalties = true
	resp.Penalty = penalty{
		PenaltyType: p.PenaltyType,
		Cause:       p.Cause,
		EndDate:     p.EndDate,
	}
}

func (resp Response) SetBooking(b entities.Booking) {
	resp.Booking = booking{
		CommonSpaceName: b.CommonSpace,
		Date:            b.Date,
		Shift:           b.Shift,
	}
}

type booking struct {
	CommonSpaceName string          `json:"common_space_name"`
	Date            time.Time       `json:"date"`
	Shift           constants.Shift `json:"shift"`
}

type penalty struct {
	PenaltyType string    `json:"type"`
	Cause       string    `json:"cause"`
	EndDate     time.Time `json:"end_date"`
}
