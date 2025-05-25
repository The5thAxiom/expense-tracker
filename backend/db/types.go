package db

import (
	"database/sql"
	"time"
)

type idNameDesc struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Category idNameDesc

type SubCategory idNameDesc

type Purpose idNameDesc

type Tag idNameDesc

type Currency struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Expense struct {
	Id          string      `json:"id"`
	Date        time.Time   `json:"date"`
	Description string      `json:"description"`
	Amount      int         `json:"amount"`
	Currency    Currency    `json:"currency"`
	Category    Category    `json:"category"`
	SubCategory SubCategory `json:"subCategory"`
	Purpose     *Purpose    `json:"purpose"`
	Notes       *string     `json:"notes"`
	Tags        []Tag       `json:"tags"`
}

type CategoryEntity struct {
	Id          string
	UserId      string
	Name        string
	Description sql.NullString
}

type SubCategoryEntity struct {
	Id          string
	UserId      string
	CategoryId  string
	Name        string
	Description sql.NullString
}

type CurrencyEntity struct {
	Id     string
	Name   string
	Symbol string
}

type PurposeEntity struct {
	Id          string
	UserId      string
	Name        string
	Description sql.NullString
}

type TagEntity struct {
	Id          string
	UserId      string
	Name        string
	Description sql.NullString
}

type ExpenseTagEntity struct {
	UserId    string
	TagId     string
	ExpenseId string
}

type ExpenseEntity struct {
	Id            string
	UserId        string
	Date          string
	Description   string
	Amount        int
	CurrencyId    string
	SubCategoryId string
	PurposeId     sql.NullString
	Notes         sql.NullString
}
