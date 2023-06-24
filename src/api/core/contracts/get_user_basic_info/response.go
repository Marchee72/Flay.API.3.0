package get_user_basic_info

import (
	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/contracts/common"
	"flay-api-v3.0/src/api/core/entities"
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	User      common.UserLw   `json:"user"`
	Building  common.Building `json:"building"`
	Apartment apartment       `json:"apartment"`
}

type apartment struct {
	ID    primitive.ObjectID `json:"_id"`
	Floor uint8              `json:"floor"`
	Flat  string             `json:"flat"`
}

func newApartment(ap entities.Apartment) apartment {
	return apartment{
		ID:    ap.ID,
		Floor: ap.Floor,
		Flat:  ap.Flat,
	}
}

func (resp *Response) NewResponse(user lw.UserLw, building *entities.Building, apartment *entities.Apartment) {
	resp.User = common.NewUserLw(user)
	if user.Type != constants.UserAdmin {
		resp.Building = common.NewBuilding(*building)
		resp.Apartment = newApartment(*apartment)
	}
}
