package book_common_space

import (
	"time"
)

type Response struct {
	Booking            booking `json:"booking"`
	HasActivePenalties bool    `json:"has_active_penalties"`
	Penalty
}

type booking struct {
	CommonSpaceName string    `json:"common_space_name"`
	InitDate        time.Time `json:"init_date"`
	FinishDate      time.Time `json:"finish_date"`
}

type penalty struct {
}
