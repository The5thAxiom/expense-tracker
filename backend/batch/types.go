package batch

import (
	"encoding/json"
	"time"
)

type ExcelExpenseRow struct {
	Date         time.Time `json:"date"`
	ExpenseIndex int       `json:"expense_index"`
	Description  string    `json:"description"`
	Amount       float64   `json:"amount"`
	Currency     string    `json:"currency"`
	Category     string    `json:"category"`
	SubCategory  string    `json:"sub_category"`
	Purpose      *string   `json:"purpose,omitempty"`
	Notes        *string   `json:"notes,omitempty"`
}

func (p ExcelExpenseRow) ToString() string {
	jsonBytes, _ := json.Marshal(p)
	return string(jsonBytes)
}
