package database

import "database/sql"

type DB interface {
	Db() *sql.DB
	Close()

	// data access methods
	GetAllCategories() ([]Category, error)
	GetCategoryById(id string) (*Category, error)
	// GetAllSubCategories() ([]SubCategory, error)
	// GetSubCategoryById(id string) (SubCategory, error)
	// GetAllCurrencies() ([]Currency, error)
	// GetCurrencyByAbbreviation(abbreviation string) (*Currency, error)
	// GetAllPurposes() ([]Purpose, error)
	// GetPurposeById(id string) (*Purpose, error)
	GetAllPayments() ([]Payment, error)
	GetPaymentById(id int) (*Payment, error)

	// data insertion methods

	// data update methods
}
