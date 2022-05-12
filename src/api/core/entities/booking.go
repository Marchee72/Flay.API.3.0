package entities

import (
	"time"

	"flay-api-v3.0/src/api/core/entities/lw"
)

type Booking struct {
	User        lw.UserLw        `bson:"user"`
	Building    lw.BuildingLw    `bson:"building"`
	CommonSpace lw.CommonSpaceLw `bson:"common_space"`
	StartDate   time.Time        `bson:"start_date"`
	EndDate     time.Time        `bson:"end_date"`
}
