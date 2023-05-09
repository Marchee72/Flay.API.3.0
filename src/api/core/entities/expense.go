package entities

import (
	"time"

	"flay-api-v3.0/src/api/core/entities/lw"
)

type Expense struct {
	Filename string        `bson:"filename"`
	Unit     Unit          `bson:"unit"`
	Building lw.BuildingLw `bson:"building"`
	Date     time.Time     `bson:"date"`
	Month    time.Month    `bson:"month"`
	Year     int           `bson:"year"`
	File     []byte
}

// func (expense *Expense) SetFile(file []byte) {
// 	expense.File = bson.RawValue{
// 		Type:  bson.TypeBinary,
// 		Value: file,
// 	}
// }

func NewExpense(name string, building lw.BuildingLw, unit Unit, month time.Month, year int, file []byte) Expense {
	expense := Expense{
		Filename: name,
		Building: building,
		Unit:     unit,
		Date:     time.Now(),
		Month:    month,
		Year:     year,
		File:     file,
	}
	//expense.SetFile(file)
	return expense
}
