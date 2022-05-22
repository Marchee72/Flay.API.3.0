package book_common_space

import (
	"time"

	"flay-api-v3.0/src/api/core/contracts/common"
)

type Request struct {
	User        common.UserLw     `json:"user"`
	CommonSpace string            `json:"common_space"`
	Builging    common.BuildingLw `json:"building"`
	StartDate   time.Time         `json:"start_date"`
	EndDate     time.Time         `json:"end_date"`
}
