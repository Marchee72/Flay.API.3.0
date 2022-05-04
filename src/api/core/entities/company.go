package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	ID             primitive.ObjectID `bson:"ID"`
	Name           string             `bson:"company_name"`
	Email          string             `bson:"company_email"`
	CompanyManager string             `bson:"company_manager"`
	Address        Address            `bson:"company_address"`
}
