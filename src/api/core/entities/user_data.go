package entities

import "flay-api-v3.0/src/api/core/entities/lw"

type UserData struct {
	Building  lw.BuildingLw
	Apartment string
}
