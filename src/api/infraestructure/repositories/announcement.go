package store

import (
	"context"
	"time"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repository *AnnouncementRepository) GetBuildingAnnouncements(ctx context.Context, buildingID primitive.ObjectID, from *time.Time, date *time.Time) ([]entities.Announcement, error) {
	var filter bson.M
	//TODO: Fix $gte date query
	if date != nil {
		//dateBson := primitive.NewDateTimeFromTime(*date)
		filter = bson.M{"building._id": buildingID /*, "date": dateBson*/}
	} else {
		//fromBson := primitive.NewDateTimeFromTime(*from)
		filter = bson.M{"building._id": buildingID /*, "date": bson.M{"$gte": fromBson}*/}
	}
	//fromBson := primitive.NewDateTimeFromTime(from)
	cursor, err := repository.Announcements.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var result []entities.Announcement
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
