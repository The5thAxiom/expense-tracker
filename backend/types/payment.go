package backend

import (
	"encoding/json"
	"log"
	"time"
)

type Category string

type SubCategory string

type Payment struct {
	Date         time.Time   `json:"date"`         // when was the payment made (only the date is relevant, the time isn't)
	PaymentIndex int         `json:"paymentIndex"` // sequence of which payment of the day it is
	Description  string      `json:"description"`
	Amount       float64     `json:"amount"` // in INR
	Currency     string      `json:"currency"`
	Category     Category    `json:"category"`
	SubCategory  SubCategory `json:"subCategory"`
	Purpose      *string     `json:"purpose"`
	Notes        *string     `json:"notes"`
}

func (p Payment) ToString() string {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonBytes)
}
