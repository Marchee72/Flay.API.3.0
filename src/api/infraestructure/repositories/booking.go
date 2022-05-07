package store

import "go.mongodb.org/mongo-driver/mongo"

type BookingRepository struct {
	DBClient *mongo.Database
}

func (repository *BookingRepository) Book() {

}

func (repository *BookingRepository) IsAbailable() (bool, error) {
	return true, nil
}
