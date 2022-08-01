package get_user_bookings

import "go.mongodb.org/mongo-driver/bson/primitive"

type Request struct {
	UserID primitive.ObjectID `json:"user_id" binding:"required"`
}
