package save_announcement

import (
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	BuildingID primitive.ObjectID `json:"building_id"`
	Message    string             `json:"message"`
	Date       time.Time          `json:"date"`
	Severity   constants.Severity `json:"severity"`
}
