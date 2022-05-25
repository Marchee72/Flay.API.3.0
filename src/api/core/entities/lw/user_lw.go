package lw

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLw struct {
	ID       primitive.ObjectID `json:"_id"`
	Username string             `json:"name"`
}
