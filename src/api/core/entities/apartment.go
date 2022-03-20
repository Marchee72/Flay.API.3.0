package entities

type Apartment struct {
	Owner User `json:"owner" bson:"owner"`
}
