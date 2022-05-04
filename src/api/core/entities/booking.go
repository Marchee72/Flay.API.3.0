package entities

import (
	"time"

	"flay-api-v3.0/src/api/core/entities/lw"
)

type Booking struct {
	User        lw.UserLw   `bson:"user"`
	CommonSpace CommonSpace `bson:"common_space"`
	StartDate   time.Time   `bson:"start_date"`
	FinishDate  time.Time   `bson:"finish_date"`
}
