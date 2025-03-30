package batch

import (
	"encoding/json"
	"time"
)

type ExcelPaymentRow struct {
	Date         time.Time `json:"date"`
	PaymentIndex int       `json:"payment_index"`
	Description  string    `json:"description"`
	Amount       float64   `json:"amount"`
	Currency     string    `json:"currency"`
	Category     string    `json:"category"`
	SubCategory  string    `json:"sub_category"`
	Purpose      *string   `json:"purpose,omitempty"`
	Notes        *string   `json:"notes,omitempty"`
}

func (p ExcelPaymentRow) ToString() string {
	jsonBytes, _ := json.Marshal(p)
	return string(jsonBytes)
}
