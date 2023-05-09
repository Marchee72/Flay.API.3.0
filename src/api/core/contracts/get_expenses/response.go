package get_expenses

import (
	"time"

	"flay-api-v3.0/src/api/core/contracts/common"
	"flay-api-v3.0/src/api/core/entities"
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
		File:     e.File,
	}
	resp.Expenses = append(resp.Expenses, newExpense)
}

type expense struct {
	Filename string            `json:"filename"`
	Building common.BuildingLw `json:"building"`
	Unit     string            `json:"unit"`
	Date     time.Time         `json:"date"`
	Month    time.Month        `json:"month"`
	Year     int               `json:"year"`
	File     []byte            `json:"file" binding:"required"`
}
