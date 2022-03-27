package providers

import (
	"flay-api-v3.0/src/api/core/entities"
)

type Authentication interface {
	CreateToken(user entities.UserLogin) (string, error)
}
