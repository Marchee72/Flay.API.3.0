package store

import (
	"context"
	"time"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PenaltyRepository struct {
	Penalties *mongo.Collection
}

func (repository *PenaltyRepository) GetActivePenalties(ctx context.Context, userID primitive.ObjectID) (*entities.Penalty, error) {
	var result entities.Penalty
	err := repository.Penalties.FindOne(ctx, bson.M{"user_id": userID, "start_date": bson.M{"$lte": time.Now()}, "end_date": bson.M{"$gte": time.Now()}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	return &result, nil
}
