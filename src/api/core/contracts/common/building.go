package common

import (
	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Building struct {
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

func NewBuilding(building entities.Building) Building {
	newBuilding := Building{
		ID:      building.ID,
		Name:    building.Name,
		Address: building.Address,
	}
	newBuilding.SetCommonSpaces(building.CommonSpaces)
	return newBuilding
}

func (resp *Building) SetCommonSpaces(commonSpaces []entities.CommonSpace) {
	for _, cs := range commonSpaces {
		resp.CommonSpaces = append(resp.CommonSpaces, commonSpace{Name: cs.Name, Type: cs.Type})
	}
}
