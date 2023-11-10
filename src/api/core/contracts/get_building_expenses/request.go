package get_building_expenses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	BuildingRawID string      `uri:"building_id" binding:"required"`
	Year          *int        `form:"year"`
	Month         *time.Month `form:"month"`
}

func (request *Request) BuildingID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(request.BuildingRawID)
	return id
}
