package entities

import "flay-api-v3.0/src/api/core/entities/lw"

type Booking struct {
	User        lw.UserLw
	CommonSpace CommonSpace
}
