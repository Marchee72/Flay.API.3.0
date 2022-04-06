package store

import "go.mongodb.org/mongo-driver/mongo"

type CommonSpaceRepository struct {
	DBClient *mongo.Database
}

func (repository *CommonSpaceRepository) Book() {

}

func (repository *CommonSpaceRepository) IsAbailable() (bool, error) {

}
