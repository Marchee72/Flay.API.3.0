package common

import (
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ComonSpaceLw struct {
	ID   primitive.ObjectID `json:"id"`
	Name string             `json:"name"`
}

func NewCommonSpaceLw(lw lw.CommonSpaceLw) ComonSpaceLw {
	return ComonSpaceLw{
		ID:   lw.ID,
		Name: lw.Name,
	}
}
