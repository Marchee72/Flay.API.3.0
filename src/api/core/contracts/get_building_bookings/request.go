package get_building_bookings

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	BuildingID string     `uri:"building_id" binding:"required"`
	Date       *time.Time `form:"date"`
}

func (request *Request) ID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(request.BuildingID)
	return id
}
