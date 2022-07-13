package book_common_space

import (
	"flay-api-v3.0/src/api/core/constants"
)

type Request struct {
	CommonSpace string          `json:"common_space"`
	Date        string          `json:"date"`
	Shift       constants.Shift `json:"shift"`
}
