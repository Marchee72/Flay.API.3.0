package save_expense_file

import (
	"time"

	"flay-api-v3.0/src/api/core/contracts/common"
)

type Request struct {
	Filename string        `json:"filename"`
	Renter   common.UserLw `json:"user"`
	Date     time.Time     `json:"date"`
	Month    time.Month    `json:"month"`
	Year     int16         `json:"year"`
	File     []byte        `json:"file"`
}
