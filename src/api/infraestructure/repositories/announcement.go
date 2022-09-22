package store

import "go.mongodb.org/mongo-driver/mongo"

type AnnouncementRepository struct {
	Announcements *mongo.Collection
}
