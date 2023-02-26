package save_announcement

import (
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	Body   Body
	Params Params
}

type Body struct {
	//BuildingID primitive.ObjectID `json:"building_id"`
	Title    string             `json:"title"`
	Message  string             `json:"message"`
	Date     time.Time          `json:"date"`
	Severity constants.Severity `json:"severity"`
}

type Params struct {
	RawBuildingID string `uri:"building_id"`
}

func (params *Params) BuildingID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(params.RawBuildingID)
	return id
}
