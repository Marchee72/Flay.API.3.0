package providers

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
)

type AnnouncementRepository interface {
	Save(ctx context.Context, announcement entities.Announcement) error
}
