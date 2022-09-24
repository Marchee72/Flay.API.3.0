package get_building_announcements

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	BuildingRawID string     `uri:"building_id" binding:"required"`
	Date          *time.Time `form:"date"`
}

func (request *Request) BuildingID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(request.BuildingRawID)
	return id
}
