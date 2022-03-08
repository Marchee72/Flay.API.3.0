package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	ID             primitive.ObjectID `json:"ID" bson:"ID"`
	Name           string             `json:"company_name" bson:"company_name"`
	Email          string             `json:"company_email" bson:"company_email"`
	CompanyManager string             `json:"company_manager" bson:"company_manager"`
	Address        Address            `json:"company_address" bson:"company_address"`
}
