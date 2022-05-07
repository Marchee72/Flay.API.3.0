package store

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IssueRepository struct {
	Issues *mongo.Collection
}

func (repository *IssueRepository) GetUserIssues(ctx context.Context, userID primitive.ObjectID) ([]entities.Issue, error) {
	var issues []entities.Issue
	cursor, err := repository.Issues.Find(ctx, bson.M{"creator.id": userID})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &issues)
	if err != mongo.ErrNoDocuments {
		return nil, err
	}
	return issues, nil
}
