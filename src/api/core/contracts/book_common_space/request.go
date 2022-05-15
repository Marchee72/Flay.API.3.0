package book_common_space

import (
	"time"

	"flay-api-v3.0/src/api/core/contracts/common"
	"flay-api-v3.0/src/api/core/entities/lw"
)

type Request struct {
	User        common.UserLw       `json:"user"`
	CommonSpace common.ComonSpaceLw `json:"common_space"`
	Builging    lw.BuildingLw       `json:"building"`
	StartDate   time.Time           `json:"start_date"`
	EndDate     time.Time           `json:"end_date"`
}
