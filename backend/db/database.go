package db

import "database/sql"

type DB interface {
	DbConn() *sql.DB

	// GetCategoryById(userId string, id string) (*Category, error)
	// GetSubCategoryForCategoryById(userId string, id string, categoryId string) (*SubCategory, error)

	GetAllExpenses(userId string) ([]Expense, error)

	InsertCategory(categeory CategoryEntity) error
	InsertTag(tag TagEntity) error
	LinkTagToExpense(tagId string, expenseId string) error
	InsertCurrency(currency CurrencyEntity) error
	InsertPurpose(purpose PurposeEntity) error
	InsertSubCategory(subCategory SubCategoryEntity) error
	InsertExpense(userId string, expense ExpenseEntity) error
}
