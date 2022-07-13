package entities

import (
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/entities/lw"
)

type Booking struct {
	User        lw.UserLw       `bson:"user"`
	Building    lw.BuildingLw   `bson:"building"`
	CommonSpace string          `bson:"common_space"`
	Date        time.Time       `bson:"date"`
	Shift       constants.Shift `bson:"shift"`
}
