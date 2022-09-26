package entities

import (
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/entities/lw"
)

type Announcement struct {
	//ID       primitive.ObjectID `bson:"_id"`
	User     lw.UserLw          `bson:"user"`
	Building lw.BuildingLw      `bson:"building"`
	Message  string             `bson:"message"`
	Date     time.Time          `bson:"date"`
	Severity constants.Severity `bson:"severity"`
}
