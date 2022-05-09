package providers

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IssueRepository interface {
	GetUserIssues(ctx context.Context, userID primitive.ObjectID) ([]entities.Issue, error)
	GetBuildingIssues(ctx context.Context, buildingID primitive.ObjectID) ([]entities.Issue, error)
}
