package entities

type Apartment struct {
	Owner User `bson:"owner"`
}
