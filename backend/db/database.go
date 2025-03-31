package db

import "database/sql"

type DB interface {
	DbConn() *sql.DB

	// data access methods
	GetAllCategories() ([]Category, error)
	GetCategoryById(id string) (*Category, error)
	GetAllSubCategoriesforCategory(categoryId string) ([]SubCategory, error)
	GetSubCategoryForCategoryById(id string, categoryId string) (*SubCategory, error)
	GetAllCurrencies() ([]Currency, error)
	GetCurrencyByAbbreviation(abbreviation string) (*Currency, error)
	GetAllPurposes() ([]Purpose, error)
	GetPurposeById(id string) (*Purpose, error)
	GetAllPayments() ([]Payment, error)
	GetPaymentById(id int) (*Payment, error)

	// data insertion methods

	// data update methods
}
