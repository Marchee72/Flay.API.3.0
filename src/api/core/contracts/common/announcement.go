package common

import (
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Announcement struct {
	ID       primitive.ObjectID `json:"id"`
	User     UserLw             `json:"user"`
	Building BuildingLw         `json:"building"`
	Title    string             `json:"title"`
	Message  string             `json:"message"`
	Date     time.Time          `json:"date"`
	Severity constants.Severity `json:"severity"`
}
