package providers

import (
	"context"
	"time"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnnouncementRepository interface {
	Save(ctx context.Context, announcement entities.Announcement) error
	GetBuildingAnnouncements(ctx context.Context, buildingID primitive.ObjectID, from *time.Time, date *time.Time) ([]entities.Announcement, error)
	GetAnnouncement(ctx context.Context, announcementID primitive.ObjectID) (*entities.Announcement, error)
}
