package batch

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

func WritePayment(db *sql.DB, payment ExcelPaymentRow) error {
	categoryId, err := getExistingOrNewCategoryId(db, payment.Category)
	if err != nil {
		return err
	}

	subCategoryId, err := getExistingOrNewSubCategoryId(db, payment.SubCategory, categoryId)
	if err != nil {
		return err
	}

	purposeId, err := getExistingOrNewPurposeId(db, payment.Purpose)
	if err != nil {
		return err
	}

	err = insertCurrencyIfDoesNotExist(db, payment.Currency)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO Payment (
			date,
			paymentIndex,
			description,
			amount,
			currencyAbbreviation,
			subCategoryId,
			purposeId,
			notes
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`,
		payment.Date,
		payment.PaymentIndex,
		payment.Description,
		payment.Amount,
		payment.Currency,
		subCategoryId,
		purposeId,
		payment.Notes,
	)
	if err != nil {
		return err
	}

	return nil
}

func getExistingOrNewCategoryId(db *sql.DB, categoryName string) (string, error) {
	categoryId := strings.ReplaceAll(categoryName, " ", "-")
	categoryId = strings.TrimSpace(categoryId)
	categoryId = strings.ToLower(categoryId)

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
	subCategoryId := strings.ReplaceAll(subCategoryName, " ", "-")
	subCategoryId = strings.TrimSpace(subCategoryId)
	subCategoryId = strings.ToLower(subCategoryId)

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

	purposeId := strings.ReplaceAll(purposeName, " ", "-")
	purposeId = strings.TrimSpace(purposeId)
	purposeId = strings.ToLower(purposeId)

	err := db.QueryRow(`SELECT id FROM Purpose WHERE id=?;`, purposeId).Scan(&purposeId)
	if err == sql.ErrNoRows {
		log.Printf("Inserting new Purpose {id: %s, name: %s}", purposeId, purposeName)
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
	err := db.QueryRow(`SELECT abbreviation FROM Currency WHERE abbreviation=?;`, currencyAbbr).Scan(&currencyAbbr)
	if err == sql.ErrNoRows {
		log.Printf("Inserting new Currency {id: %s, name: %s}", currencyAbbr, currencyAbbr)
		_, err = db.Exec(`INSERT INTO Currency (abbreviation, name) VALUES (?, ?);`, currencyAbbr, currencyAbbr)
		if err != nil {
			return errors.New("Error inserting into Currency table: " + err.Error())
		}
	}
	if err != nil {
		return errors.New("Error querying Currency table: " + err.Error())
	}

	return nil
}
