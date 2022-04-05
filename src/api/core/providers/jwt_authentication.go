package providers

import (
	"flay-api-v3.0/src/api/core/entities/lw"
)

type Authentication interface {
	CreateToken(user lw.UserLw) (string, error)
}
