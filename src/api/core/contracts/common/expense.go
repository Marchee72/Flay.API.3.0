package common

import (
	"mime/multipart"
	"time"
)

type Expense struct {
	Filename string                `form:"filename"`
	Building BuildingLw            `form:"building"`
	Unit     string                `form:"unit"`
	Date     time.Time             `form:"date"`
	Month    time.Month            `form:"month"`
	Year     int                   `form:"year"`
	File     *multipart.FileHeader `form:"file" binding:"required"`
}
