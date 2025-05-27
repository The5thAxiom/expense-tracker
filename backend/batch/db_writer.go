package batch

import (
	"backend/db"
	"database/sql"
	"errors"
	"log"
	"strings"
)

func WriteExpense(d db.DB, expense ExcelExpenseRow, index int) error {
	db := d.DbConn()
	log.Printf("Writing expense #%d (%s #%d)\n", index, expense.Date.String(), expense.ExpenseIndex)

	categoryId, err := getExistingOrNewCategoryId(db, expense.Category)
	if err != nil {
		return err
	}

	subCategoryId, err := getExistingOrNewSubCategoryId(db, expense.SubCategory, categoryId)
	if err != nil {
		return err
	}

	purposeId, err := getExistingOrNewPurposeId(db, expense.Purpose)
	if err != nil {
		return err
	}

	err = insertCurrencyIfDoesNotExist(db, expense.Currency)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO Expense (
			date,
			expenseIndex,
			description,
			amount,
			currencyId,
			subCategoryId,
			purposeId,
			notes
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`,
		expense.Date,
		expense.ExpenseIndex,
		expense.Description,
		expense.Amount,
		expense.Currency,
		subCategoryId,
		purposeId,
		expense.Notes,
	)
	if err != nil {
		return err
	}

	return nil
}

func getExistingOrNewCategoryId(db *sql.DB, categoryName string) (string, error) {
	categoryId := makeIdFromString(categoryName)

	err := db.QueryRow(`SELECT id FROM Category WHERE id=?;`, categoryId).Scan(&categoryId)
	if err == sql.ErrNoRows {
		log.Printf("Inserting new Category {id: %s, name: %s}", categoryId, categoryName)
		_, err = db.Exec(`INSERT INTO Category (id, name) VALUES (?, ?);`, categoryId, categoryName)
		if err != nil {
			return "", errors.New("Error inserting into Category table: " + err.Error())
		}
	}
	if err != nil {
		return "", errors.New("Error querying Category table: " + err.Error())
	}

	return categoryId, nil
}

func getExistingOrNewSubCategoryId(db *sql.DB, subCategoryName string, categoryId string) (string, error) {
	subCategoryId := makeIdFromString(subCategoryName)

	err := db.QueryRow(`SELECT id FROM SubCategory WHERE id=? AND categoryId=?;`, subCategoryId, categoryId).Scan(&subCategoryId)
	if err == sql.ErrNoRows {
		log.Printf("Inserting new SubCategory {id: %s, name: %s} for category %s", subCategoryId, subCategoryName, categoryId)
		_, err = db.Exec(`INSERT INTO SubCategory (id, name, categoryId) VALUES (?, ?, ?);`, subCategoryId, subCategoryName, categoryId)
		if err != nil {
			return "", errors.New("Error inserting into SubCategory table: " + err.Error())
		}
	}
	if err != nil {
		return "", errors.New("Error querying SubCategory table: " + err.Error())
	}

	return subCategoryId, nil
}

func getExistingOrNewPurposeId(db *sql.DB, purpose *string) (*string, error) {
	if purpose == nil {
		return nil, nil
	}

	purposeName := *purpose
	purposeId := makeIdFromString(purposeName)

	err := db.QueryRow(`SELECT id FROM Purpose WHERE id=?;`, purposeId).Scan(&purposeId)
	if err == sql.ErrNoRows {
		log.Printf("Inserting new Purpose {id: '%s', name: '%s'}", purposeId, purposeName)
		_, err = db.Exec(`INSERT INTO Purpose (id, name) VALUES (?, ?);`, purposeId, purposeName)
		if err != nil {
			return nil, errors.New("Error inserting into Purpose table: " + err.Error())
		}
	}
	if err != nil {
		return nil, errors.New("Error querying Purpose table: " + err.Error())
	}

	return &purposeId, nil
}

func insertCurrencyIfDoesNotExist(db *sql.DB, currencyAbbr string) error {
	err := db.QueryRow(`SELECT id FROM Currency WHERE id=?;`, currencyAbbr).Scan(&currencyAbbr)
	if err == sql.ErrNoRows {
		log.Printf("Inserting new Currency {id: %s, name: %s}", currencyAbbr, currencyAbbr)
		_, err = db.Exec(`INSERT INTO Currency (id, name) VALUES (?, ?);`, currencyAbbr, currencyAbbr)
		if err != nil {
			return errors.New("Error inserting into Currency table: " + err.Error())
		}
	}
	if err != nil {
		return errors.New("Error querying Currency table: " + err.Error())
	}

	return nil
}

func makeIdFromString(str string) string {
	id := strings.ToLower(str)
	id = strings.TrimSpace(id)
	id = strings.ReplaceAll(id, " ", "-")
	return id
}
