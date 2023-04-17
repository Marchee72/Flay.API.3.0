package entities

import (
	"time"

	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson"
)

type Expense struct {
	Filename string        `bson:"filename"`
	Renter   lw.UserLw     `bson:"user"`
	Date     time.Time     `bson:"date"`
	Month    time.Month    `bson:"month"`
	Year     int           `bson:"year"`
	Content  bson.RawValue `bson:"content"`
}

func (expense *Expense) SetFile(file []byte) {
	expense.Content = bson.RawValue{
		Type:  bson.TypeBinary,
		Value: file,
	}
}

func NewExpense(name string, renter lw.UserLw, month time.Month, year int, file []byte) Expense {
	expense := Expense{
		Filename: name,
		Renter:   renter,
		Date:     time.Now(),
		Month:    month,
		Year:     year,
	}
	expense.SetFile(file)
	return expense
}
