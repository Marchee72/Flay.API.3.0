package dependencies

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/gridfs"
)

type CollectionContainer struct {
	Users         *mongo.Collection
	CommonSpaces  *mongo.Collection
	Penalties     *mongo.Collection
	Bookings      *mongo.Collection
	Issues        *mongo.Collection
	Buildings     *mongo.Collection
	Announcements *mongo.Collection
	Apartments    *mongo.Collection
	Expenses      *gridfs.Bucket
}

func NewCollectionContainer(db *mongo.Database) CollectionContainer {
	return CollectionContainer{
		Users:         db.Collection("users"),
		CommonSpaces:  db.Collection("common_spaces"),
		Penalties:     db.Collection("penalties"),
		Bookings:      db.Collection("bookings"),
		Issues:        db.Collection("issues"),
		Buildings:     db.Collection("buildings"),
		Announcements: db.Collection("announcements"),
		Apartments:    db.Collection("apartments"),
	}

}
