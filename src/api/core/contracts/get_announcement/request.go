package get_announcement

import "go.mongodb.org/mongo-driver/bson/primitive"

type Request struct {
	AnnouncementRawID string `uri:"announcement_id" binding:"required"`
}

func (request *Request) BuildingID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(request.AnnouncementRawID)
	return id
}
