package save_announcement

import (
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/entities/lw"
)

type Request struct {
	Building lw.BuildingLw      `json:"building"`
	Message  string             `json:"message"`
	Date     time.Time          `json:"date"`
	Severity constants.Severity `json:"severity"`
}
