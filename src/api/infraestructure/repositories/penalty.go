package store

import (
	"errors"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PenaltyRepository struct {
	DBClient *mongo.Database
}

func (repository *PenaltyRepository) GetActivePenalties(userID primitive.ObjectID) (*entities.Penalty, error) {
	return nil, errors.New("")
}
