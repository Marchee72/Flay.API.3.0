package common

import (
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLw struct {
	ID   primitive.ObjectID `json:"id"`
	Name string             `json:"name"`
}

func NewUserLw(lw lw.UserLw) UserLw {
	return UserLw{
		ID:   lw.ID,
		Name: lw.Username,
	}
}
