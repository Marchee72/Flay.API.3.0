package get_building_expenses

import "flay-api-v3.0/src/api/core/contracts/common"

type Response struct {
	Expenses []common.Expense `json:"expenses"`
}
