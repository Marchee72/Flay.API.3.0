package get_user_building

import (
	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	ID           primitive.ObjectID `json:"id"`
	Name         string             `json:"name"`
	Address      string             `json:"address"`
	CommonSpaces []commonSpace      `json:"common_spaces"`
}

type address struct {
	Street    string `json:"srteet"`
	Number    uint64 `json:"number"`
	Floor     uint64 `json:"floor"`
	Apartment string `json:"apartment"`
}

type commonSpace struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (resp *Response) SetCommonSpaces(commonSpaces []entities.CommonSpace) {
	for _, cs := range commonSpaces {
		resp.CommonSpaces = append(resp.CommonSpaces, commonSpace{Name: cs.Name, Type: cs.Type})
	}
}

// func (resp *Response) SetAddres(addres entities.Address) {
// 	resp.Address = address{
// 		Street:    addres.Street,
// 		Number:    addres.Number,
// 		Apartment: addres.Apartment,
// 		Floor:     addres.Floor,
// 	}
// }
