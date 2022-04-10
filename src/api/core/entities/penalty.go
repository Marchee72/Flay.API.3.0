package entities

import (
	"time"

	"flay-api-v3.0/src/api/core/entities/lw"
)

type Penalty struct {
	User        lw.UserLw
	PenaltyType string
	Cause       string
	Amount      int64
	EndDate     time.Time
}
