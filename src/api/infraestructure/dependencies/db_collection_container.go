package dependencies

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CollectionContainer struct {
	Files         *mongo.Collection
	Users         *mongo.Collection
	CommonSpaces  *mongo.Collection
	Penalties     *mongo.Collection
	Bookings      *mongo.Collection
	Issues        *mongo.Collection
	Buildings     *mongo.Collection
	Announcements *mongo.Collection
	Apartments    *mongo.Collection
	Expenses      *mongo.Collection
}

func NewCollectionContainer(db *mongo.Database) CollectionContainer {
	return CollectionContainer{
		Files:         db.Collection("files"),
		Users:         db.Collection("users"),
		CommonSpaces:  db.Collection("common_spaces"),
		Penalties:     db.Collection("penalties"),
		Bookings:      db.Collection("bookings"),
		Issues:        db.Collection("issues"),
		Buildings:     db.Collection("buildings"),
		Announcements: db.Collection("announcements"),
		Apartments:    db.Collection("apartments"),
		Expenses:      db.Collection("expenses"),
	}

}
