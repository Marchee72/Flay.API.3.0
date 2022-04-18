package book_common_space

import (
	"time"

	"flay-api-v3.0/src/api/core/entities"
)

type Response struct {
	Booking            booking `json:"booking"`
	HasActivePenalties bool    `json:"has_active_penalties"`
	Penalty            penalty `json:"penalty"`
}

func (resp Response) SetPenalty(p entities.Penalty) {
	resp.HasActivePenalties = true
	resp.Penalty = penalty{
		PenaltyType: p.PenaltyType,
		Cause:       p.Cause,
		EndDate:     p.EndDate,
	}
}

type booking struct {
	CommonSpaceName string    `json:"common_space_name"`
	InitDate        time.Time `json:"init_date"`
	FinishDate      time.Time `json:"finish_date"`
}

type penalty struct {
	PenaltyType string
	Cause       string
	EndDate     time.Time
}
