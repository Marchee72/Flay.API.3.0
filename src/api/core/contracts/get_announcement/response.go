package get_announcement

import (
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/contracts/common"
)

type Response struct {
	User     common.UserLw      `json:"user"`
	Building common.BuildingLw  `json:"building"`
	Title    string             `json:"title"`
	Message  string             `json:"message"`
	Date     time.Time          `json:"date"`
	Severity constants.Severity `json:"severity"`
}
