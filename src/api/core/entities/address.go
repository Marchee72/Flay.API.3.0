package entities

type Address struct {
	Street    string `json:"srteet" bson:"srteet"`
	Number    uint64 `json:"number" bson:"number"`
	Floor     uint64 `json:"floor" bson:"floor"`
	Apartment string `json:"apartment" bson:"apartment"`
}
