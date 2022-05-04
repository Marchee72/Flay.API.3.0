package entities

import (
	"time"

	"flay-api-v3.0/src/api/core/entities/lw"
)

type Penalty struct {
	User        lw.UserLw `bson:"user"`
	PenaltyType string    `bson:"penalty_type"`
	Cause       string    `bson:"cause"`
	Amount      int64     `bson:"amount"`
	StartDate   time.Time `bson:"start_date"`
	EndDate     time.Time `bson:"end_date"`
}
