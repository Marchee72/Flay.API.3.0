package dependencies

import "go.mongodb.org/mongo-driver/mongo"

type CollectionContainer struct {
	Users        *mongo.Collection
	CommonSpaces *mongo.Collection
	Penalties    *mongo.Collection
	Bookings     *mongo.Collection
	Issues       *mongo.Collection
	Building     *mongo.Collection
}

func NewCollectionContainer(db *mongo.Database) CollectionContainer {
	return CollectionContainer{
		Users:        db.Collection("users"),
		CommonSpaces: db.Collection("common_spaces"),
		Penalties:    db.Collection("penalties"),
		Bookings:     db.Collection("bookings"),
		Issues:       db.Collection("issues"),
		Building:     db.Collection("building"),
	}

}
