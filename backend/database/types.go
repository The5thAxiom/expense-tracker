package database

import (
	"time"
)

type Category struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type SubCategory struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Purpose struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Currency struct {
	Abbreviation string  `json:"abbreviation"`
	Name         *string `json:"name"`
	Symbol       *string `json:"symbol"`
}

type Payment struct {
	Id           int         `json:"id"`
	Date         time.Time   `json:"date"`
	PaymentIndex int         `json:"paymentIndex"`
	Description  string      `json:"description"`
	Amount       float64     `json:"amount"`
	Currency     Currency    `json:"currency"`
	Category     Category    `json:"category"`
	SubCategory  SubCategory `json:"subCategory"`
	Purpose      *Purpose    `json:"purpose"`
	Notes        *string     `json:"notes"`
}
