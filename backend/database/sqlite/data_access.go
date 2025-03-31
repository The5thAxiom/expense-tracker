package sqlite

import (
	"backend/database"
	"database/sql"
	"log"
	"time"
)

func (d SQLiteDB) GetAllCategories() ([]database.Category, error) {
	rows, err := d.db.Query(`SELECT id, name, description FROM Category;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]database.Category, 0)

	for rows.Next() {
		var category database.Category
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

func (d SQLiteDB) GetCategoryById(categoryId string) (*database.Category, error) {
	var category database.Category
	var description sql.NullString

	err := d.db.QueryRow(
		`SELECT id, name, description FROM CATEGORY WHERE id=?;`, categoryId,
	).Scan(&category.Id, &category.Name, category.Description)
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

// func (d SQLiteDB) GetAllSubCategories() ([]database.SubCategory, error) {}

// func (d SQLiteDB) GetSubCategoryById(id string) (*database.SubCategory, error) {}

// func (d SQLiteDB) GetAllCurrencies() ([]database.Currency, error) {}

// func (d SQLiteDB) GetCurrencyByAbbreviation(abbreviation string) (*database.Currency, error) {}

// func (d SQLiteDB) GetAllPurposes() ([]database.Purpose, error) {}

// func (d SQLiteDB) GetPurposeById(id string) (*database.Purpose, error) {}

func (d SQLiteDB) GetAllPayments() ([]database.Payment, error) {
	rows, err := d.db.Query(`
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

	payments := make([]database.Payment, 0)

	for rows.Next() {
		var payment database.Payment

		var paymentDate string
		var paymentNotes sql.NullString

		var currencyName sql.NullString
		var currencySymbol sql.NullString

		var categoryDescription sql.NullString

		var subCategoryDescription sql.NullString

		var purposeId sql.NullString
		var purposeName sql.NullString
		var purposeDescription sql.NullString

		err = rows.Scan(
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
			return payments, err
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
			var purpose database.Purpose
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

		payments = append(payments, payment)
	}

	return payments, nil
}

func (d SQLiteDB) GetPaymentById(id int) (*database.Payment, error) {
	var payment database.Payment

	var paymentDate string
	var paymentNotes sql.NullString

	var currencyName sql.NullString
	var currencySymbol sql.NullString

	var categoryDescription sql.NullString

	var subCategoryDescription sql.NullString

	var purposeId sql.NullString
	var purposeName sql.NullString
	var purposeDescription sql.NullString

	err := d.db.QueryRow(`
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
	`, id).Scan(
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
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
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
		var purpose database.Purpose
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

	return &payment, nil
}
