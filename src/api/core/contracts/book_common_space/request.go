package book_common_space

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	CommonSpaceID primitive.ObjectID `json:"common_space_id"`
	UserID        primitive.ObjectID `json:"user_id"`
	InitDate      time.Time          `json:"init_date"`
	FinishDate    time.Time          `json:"finish_date"`
}
