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
	filter := bson.M{
		"creator.id": userID,
	}
	issues, err := repository.getIssues(ctx, filter)
	if err != nil {
		return nil, err
	}
	return issues, nil
}

func (repository *IssueRepository) GetBuildingIssues(ctx context.Context, buildingID primitive.ObjectID) ([]entities.Issue, error) {
	filter := bson.M{
		"building.id": buildingID,
	}
	issues, err := repository.getIssues(ctx, filter)
	if err != nil {
		return nil, err
	}
	return issues, nil
}

func (repository *IssueRepository) getIssues(ctx context.Context, filter bson.M) ([]entities.Issue, error) {
	cursor, err := repository.Issues.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var issues []entities.Issue
	if err = cursor.All(ctx, &issues); err != mongo.ErrNoDocuments {
		return nil, err
	}
	return issues, nil
}
