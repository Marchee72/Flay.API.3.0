package entities

import (
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/entities/lw"
)

type Issue struct {
	Creator     lw.UserLw          `bson:"creator"`
	Building    lw.BuildingLw      `bson:"building"`
	Date        time.Time          `bson:"date"`
	Status      string             `bson:"status"`
	Category    string             `bson:"category"`
	Descripton  string             `bson:"description"`
	Severity    constants.Severity `bson:"severity"`
	CommonSpace bool               `bson:"common_space"`
}
