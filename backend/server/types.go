package server

import (
	"time"
)

type MaybeNew struct {
	IsNew       bool    `json:"isNew"`
	Id          *string `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type NewExpense struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Notes       *string   `json:"notes"`

	CurrencyId string `json:"currencyId"`

	Category MaybeNew `json:"category"`

	SubCategory MaybeNew `json:"subCategory"`

	Purpose *MaybeNew `json:"purpose"`

	Tags []MaybeNew `json:"tags"`
}

type Id struct {
	Id string `json:"id"`
}
