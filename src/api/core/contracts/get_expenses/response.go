package get_expenses

import (
	"time"

	"flay-api-v3.0/src/api/core/contracts/common"
	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Expenses []expense `json:"expenses"`
}

func (resp *Response) AddExpense(e entities.Expense) {
	newExpense := expense{
		Filename: e.Filename,
		Building: common.NewBuildingLw(e.Building),
		Unit:     e.Unit.String(),
		Date:     e.Date,
		Month:    e.Month,
		Year:     e.Year,
	}
	resp.Expenses = append(resp.Expenses, newExpense)
}

type expense struct {
	ID       primitive.ObjectID `json:"id"`
	Filename string             `json:"filename"`
	Building common.BuildingLw  `json:"building"`
	Unit     string             `json:"unit"`
	Date     time.Time          `json:"date"`
	Month    time.Month         `json:"month"`
	Year     int                `json:"year"`
}
