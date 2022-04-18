package providers

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PenaltyRepository interface {
	GetActivePenalties(ctx context.Context, userID primitive.ObjectID) (*entities.Penalty, error)
}
