package lw

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommonSpaceLw struct {
	ID   primitive.ObjectID `json:"_id"`
	Name string             `json:"name"`
}
