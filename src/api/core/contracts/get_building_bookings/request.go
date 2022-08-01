package get_building_bookings

import "go.mongodb.org/mongo-driver/bson/primitive"

type Request struct {
	BuildingID primitive.ObjectID `json:"building_id" binding:"required"`
}
