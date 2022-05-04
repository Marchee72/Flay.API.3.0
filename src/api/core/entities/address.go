package entities

type Address struct {
	Street    string `bson:"srteet"`
	Number    uint64 `bson:"number"`
	Floor     uint64 `bson:"floor"`
	Apartment string `bson:"apartment"`
}
