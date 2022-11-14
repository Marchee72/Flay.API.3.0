package get_user_bookings

import "go.mongodb.org/mongo-driver/bson/primitive"

type Request struct {
	UserRawID string `uri:"user_id" binding:"required"`
}

func (request *Request) UserID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(request.UserRawID)
	return id
}
