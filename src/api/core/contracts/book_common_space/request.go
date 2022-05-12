package book_common_space

import (
	"time"

	"flay-api-v3.0/src/api/core/entities/lw"
)

type Request struct {
	User        lw.UserLw        `json:"user"`
	CommonSpace lw.CommonSpaceLw `json:"common_space_id"`
	Builging    lw.BuildingLw    `json:"building"`
	StartDate   time.Time        `json:"start_date"`
	EndDate     time.Time        `json:"end_date"`
}
