package entities

import "flay-api-v3.0/src/api/core/entities/lw"

type Penalties struct {
	User        lw.UserLw
	PenaltyType string
	Cause       string
}
