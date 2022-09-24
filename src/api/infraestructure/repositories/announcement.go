package store

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnnouncementRepository struct {
	Announcements *mongo.Collection
}

func (repository *AnnouncementRepository) Save(ctx context.Context, announcement entities.Announcement) error {
	var err error
	_, err = repository.Announcements.InsertOne(ctx, announcement)
	return err
}
