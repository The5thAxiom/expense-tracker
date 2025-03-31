package sqlite

import (
	"backend/db"
	"database/sql"
	"log"
	"time"
)

func (d SQLiteDB) GetAllCategories() ([]db.Category, error) {
	rows, err := d.dbConn.Query(`SELECT id, name, description FROM Category;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]db.Category, 0)

	for rows.Next() {
		var category db.Category
		var description sql.NullString

		err := rows.Scan(&category.Id, &category.Name, &description)
		if err != nil {
			return categories, err
		}

		if description.Valid {
			category.Description = &description.String
		} else {
			category.Description = nil
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (d SQLiteDB) GetCategoryById(id string) (*db.Category, error) {
	var category db.Category
	var description sql.NullString

	err := d.dbConn.QueryRow(
		`SELECT id, name, description FROM Category WHERE id=?;`, id,
	).Scan(&category.Id, &category.Name, &description)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return &category, err
	}

	if description.Valid {
		category.Description = &description.String
	} else {
		category.Description = nil
	}

	return &category, nil
}

func (d SQLiteDB) GetAllSubCategoriesforCategory(categoryId string) ([]db.SubCategory, error) {
	rows, err := d.dbConn.Query(`SELECT id, name, description FROM SubCategory WHERE SubCategory.categoryId=?;`, categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subCategories := make([]db.SubCategory, 0)

	for rows.Next() {
		var subCategory db.SubCategory
		var description sql.NullString

		err := rows.Scan(&subCategory.Id, &subCategory.Name, &description)
		if err != nil {
			return subCategories, err
		}

		if description.Valid {
			subCategory.Description = &description.String
		} else {
			subCategory.Description = nil
		}

		subCategories = append(subCategories, subCategory)
	}

	return subCategories, nil
}

func (d SQLiteDB) GetSubCategoryForCategoryById(id string, categoryId string) (*db.SubCategory, error) {
	var subCategory db.SubCategory
	var description sql.NullString

	err := d.dbConn.QueryRow(
		`SELECT id, name, description FROM SubCategory WHERE id=? AND categoryId=?;`, id, categoryId,
	).Scan(&subCategory.Id, &subCategory.Name, &description)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return &subCategory, err
	}

	if description.Valid {
		subCategory.Description = &description.String
	} else {
		subCategory.Description = nil
	}

	return &subCategory, nil
}

func (d SQLiteDB) GetAllCurrencies() ([]db.Currency, error) {
	rows, err := d.dbConn.Query(`SELECT abbreviation, name, symbol FROM Currency;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	currencies := make([]db.Currency, 0)

	for rows.Next() {
		var currency db.Currency
		var name sql.NullString
		var symbol sql.NullString

		err := rows.Scan(&currency.Abbreviation, &name, &symbol)
		if err != nil {
			return currencies, err
		}

		if name.Valid {
			currency.Name = &name.String
		} else {
			currency.Name = nil
		}

		if symbol.Valid {
			currency.Symbol = &symbol.String
		} else {
			currency.Symbol = nil
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (d SQLiteDB) GetCurrencyByAbbreviation(abbreviation string) (*db.Currency, error) {
	var currency db.Currency
	var name sql.NullString
	var symbol sql.NullString

	err := d.dbConn.QueryRow(
		`SELECT abbreviation, name, symbol FROM Currency WHERE abbreviation=?;`, abbreviation,
	).Scan(&currency.Abbreviation, &name, &symbol)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return &currency, err
	}

	if name.Valid {
		currency.Name = &name.String
	} else {
		currency.Name = nil
	}

	if symbol.Valid {
		currency.Symbol = &symbol.String
	} else {
		currency.Symbol = nil
	}

	return &currency, nil
}

func (d SQLiteDB) GetAllPurposes() ([]db.Purpose, error) {
	rows, err := d.dbConn.Query(`SELECT id, name, description FROM Purpose;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	purposes := make([]db.Purpose, 0)

	for rows.Next() {
		var purpose db.Purpose
		var description sql.NullString

		err := rows.Scan(&purpose.Id, &purpose.Name, &description)
		if err != nil {
			return purposes, err
		}

		if description.Valid {
			purpose.Description = &description.String
		} else {
			purpose.Description = nil
		}

		purposes = append(purposes, purpose)
	}

	return purposes, nil
}

func (d SQLiteDB) GetPurposeById(id string) (*db.Purpose, error) {
	var purpose db.Purpose
	var description sql.NullString

	err := d.dbConn.QueryRow(
		`SELECT id, name, description FROM Purpose WHERE id=?;`, id,
	).Scan(&purpose.Id, &purpose.Name, &description)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return &purpose, err
	}

	if description.Valid {
		purpose.Description = &description.String
	} else {
		purpose.Description = nil
	}

	return &purpose, nil
}

func scanPayment(scanner interface {
	Scan(dest ...interface{}) error
}) (db.Payment, error) {
	var payment db.Payment
	var paymentDate string
	var paymentNotes sql.NullString

	var currencyName sql.NullString
	var currencySymbol sql.NullString

	var categoryDescription sql.NullString

	var subCategoryDescription sql.NullString

	var purposeId sql.NullString
	var purposeName sql.NullString
	var purposeDescription sql.NullString

	err := scanner.Scan(
		&payment.Id,
		&paymentDate,
		&payment.PaymentIndex,
		&payment.Description,
		&payment.Amount,
		&paymentNotes,

		&payment.Currency.Abbreviation,
		&currencyName,
		&currencySymbol,

		&payment.Category.Id,
		&payment.Category.Name,
		&categoryDescription,

		&payment.SubCategory.Id,
		&payment.SubCategory.Name,
		&subCategoryDescription,

		&purposeId,
		&purposeName,
		&purposeDescription,
	)
	if err != nil {
		return payment, err
	}

	payment.Date, err = time.Parse("2006-01-02 03:04:05-07:00", paymentDate)
	if err != nil {
		log.Fatalf("Date format incorrect for %s", paymentDate)
	}

	if paymentNotes.Valid {
		payment.Notes = &paymentNotes.String
	}

	if currencyName.Valid {
		payment.Currency.Name = &currencyName.String
	}
	if currencySymbol.Valid {
		payment.Currency.Symbol = &currencySymbol.String
	}

	if categoryDescription.Valid {
		payment.Category.Description = &categoryDescription.String
	}

	if subCategoryDescription.Valid {
		payment.SubCategory.Description = &subCategoryDescription.String
	}

	if !purposeId.Valid && !purposeName.Valid && !purposeDescription.Valid {
		payment.Purpose = nil
	} else {
		var purpose db.Purpose
		if purposeId.Valid {
			purpose.Id = purposeId.String
		}
		if purposeName.Valid {
			purpose.Name = purposeName.String
		}
		if purposeDescription.Valid {
			purpose.Description = &purposeDescription.String
		}
		payment.Purpose = &purpose
	}
	return payment, nil
}

func (d SQLiteDB) GetAllPayments() ([]db.Payment, error) {
	rows, err := d.dbConn.Query(`
		SELECT
			Payment.id,
			Payment.date,
			Payment.paymentIndex,
			Payment.description,
			Payment.amount,
			Payment.notes,

			Currency.abbreviation,
			Currency.name,
			Currency.symbol,

			Category.id,
			Category.name,
			Category.description,

			SubCategory.id,
			SubCategory.name,
			SubCategory.description,

			Purpose.id,
			Purpose.name,
			Purpose.description
		FROM Payment
		LEFT JOIN Currency ON Payment.currencyAbbreviation = Currency.abbreviation
		LEFT JOIN Purpose ON Payment.purposeId = Purpose.id
		LEFT JOIN SubCategory ON Payment.subCategoryId = SubCategory.id
		LEFT JOIN Category ON SubCategory.categoryId = Category.id;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payments := make([]db.Payment, 0)

	for rows.Next() {
		payment, err := scanPayment(rows)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func (d SQLiteDB) GetPaymentById(id int) (*db.Payment, error) {
	payment, err := scanPayment(d.dbConn.QueryRow(`
		SELECT
			Payment.id,
			Payment.date,
			Payment.paymentIndex,
			Payment.description,
			Payment.amount,
			Payment.notes,

			Currency.abbreviation,
			Currency.name,
			Currency.symbol,

			Category.id,
			Category.name,
			Category.description,

			SubCategory.id,
			SubCategory.name,
			SubCategory.description,

			Purpose.id,
			Purpose.name,
			Purpose.description
		FROM Payment
		LEFT JOIN Currency ON Payment.currencyAbbreviation = Currency.abbreviation
		LEFT JOIN Purpose ON Payment.purposeId = Purpose.id
		LEFT JOIN SubCategory ON Payment.subCategoryId = SubCategory.id
		LEFT JOIN Category ON SubCategory.categoryId = Category.id
		WHERE Payment.id = ?;
	`, id))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return &payment, nil
	}

	return &payment, nil
}
