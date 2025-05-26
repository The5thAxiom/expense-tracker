package sqlite

import (
	"backend/db"
	"database/sql"
	"log"
	"time"
)

func (d SQLiteDB) GetCategoryById(userId string, id string) (*db.CategoryEntity, error) {
	var category db.CategoryEntity

	err := d.dbConn.QueryRow(
		`SELECT id, userId, name, description FROM Category WHERE userId=? AND id=?;`,
		userId, id,
	).Scan(&category.Id, &category.UserId, &category.Name, &category.Description)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (d SQLiteDB) GetCategories(userId string) ([]db.CategoryEntity, error) {
	rows, err := d.dbConn.Query(
		`SELECT id, userId, name, description FROM Category WHERE userId=?;`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]db.CategoryEntity, 0)

	for rows.Next() {
		var category db.CategoryEntity

		err := rows.Scan(&category.Id, &category.UserId, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (d SQLiteDB) GetSubCategoryForCategoryById(userId string, id string, categoryId string) (*db.SubCategory, error) {
	var subCategory db.SubCategory
	var description sql.NullString

	err := d.dbConn.QueryRow(
		`SELECT id, name, description FROM SubCategory WHERE userId=? AND id=? AND categoryId=?;`, userId, id, categoryId,
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

func scanExpense(scanner interface {
	Scan(dest ...interface{}) error
}) (db.Expense, error) {
	var expense db.Expense
	var expenseDate string
	var expenseNotes sql.NullString

	var currencyName sql.NullString
	var currencySymbol sql.NullString

	var categoryDescription sql.NullString

	var subCategoryDescription sql.NullString

	var purposeId sql.NullString
	var purposeName sql.NullString
	var purposeDescription sql.NullString

	err := scanner.Scan(
		&expense.Id,
		&expenseDate,
		&expense.Description,
		&expense.Amount,
		&expenseNotes,

		&expense.Currency.Id,
		&currencyName,
		&currencySymbol,

		&expense.Category.Id,
		&expense.Category.Name,
		&categoryDescription,

		&expense.SubCategory.Id,
		&expense.SubCategory.Name,
		&subCategoryDescription,

		&purposeId,
		&purposeName,
		&purposeDescription,
	)
	if err != nil {
		return expense, err
	}

	expense.Date, err = time.Parse("2006-01-02", expenseDate)
	if err != nil {
		// should not really be happening
		log.Fatalf("Date format incorrect for %s", expenseDate)
	}

	if expenseNotes.Valid {
		expense.Notes = &expenseNotes.String
	}

	if currencyName.Valid {
		expense.Currency.Name = currencyName.String
	}
	if currencySymbol.Valid {
		expense.Currency.Symbol = currencySymbol.String
	}

	if categoryDescription.Valid {
		expense.Category.Description = &categoryDescription.String
	}

	if subCategoryDescription.Valid {
		expense.SubCategory.Description = &subCategoryDescription.String
	}

	if !purposeId.Valid && !purposeName.Valid && !purposeDescription.Valid {
		expense.Purpose = nil
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
		expense.Purpose = &purpose
	}
	return expense, nil
}

func (d SQLiteDB) GetAllExpenses(userId string) ([]db.Expense, error) {
	rows, err := d.dbConn.Query(`
		SELECT
			Expense.id,
			Expense.date,
			Expense.description,
			Expense.amount,
			Expense.notes,

			Currency.id,
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
		FROM Expense
			LEFT JOIN Currency ON Expense.currencyId = Currency.id
			LEFT JOIN Purpose ON Expense.purposeId = Purpose.id
			LEFT JOIN SubCategory ON Expense.subCategoryId = SubCategory.id
			LEFT JOIN Category ON SubCategory.categoryId = Category.id
		WHERE Expense.userId = ?;
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	expenses := make([]db.Expense, 0)

	for rows.Next() {
		expense, err := scanExpense(rows)
		if err != nil {
			return nil, err
		}

		// for each expense, we need to collate all its tags
		var tags []db.Tag
		etRows, err := d.DbConn().Query(`
			SELECT
				Tag.id,
				Tag.name,
				Tag.description
			FROM ExpenseTag
				JOIN Tag ON ExpenseTag.tagId = Tag.id
			WHERE expenseId = ?;
		`, expense.Id)
		if err != nil {
			return nil, err
		}

		for etRows.Next() {
			var tag db.Tag
			var tagDescription sql.NullString

			err = etRows.Scan(&tag.Id, &tag.Name, &tagDescription)
			if err != nil {
				return nil, err
			}

			if tagDescription.Valid {
				tag.Description = &tagDescription.String
			}

			tags = append(tags, tag)
		}
		expense.Tags = tags

		expenses = append(expenses, expense)
	}

	return expenses, nil
}

// func (d SQLiteDB) GetExpenses(userId string) ([]db.Expense, error) {
// 	rows, err := d.dbConn.Query(`
// 		SELECT
// 			Expense.id,
// 			Expense.userId
// 			Expense.date,
// 			Expense.description,
// 			Expense.amount,
// 			Expense.notes,

// 			Currency.id,
// 			Currency.name,
// 			Currency.symbol,

// 			Category.id,
// 			Category.name,
// 			Category.description,

// 			SubCategory.id,
// 			SubCategory.name,
// 			SubCategory.description,

// 			Purpose.id,
// 			Purpose.name,
// 			Purpose.description
// 		FROM Expense
// 			LEFT JOIN Currency ON Expense.currencyId = Currency.id
// 			LEFT JOIN Purpose ON Expense.purposeId = Purpose.id
// 			LEFT JOIN SubCategory ON Expense.subCategoryId = SubCategory.id
// 			LEFT JOIN Category ON SubCategory.categoryId = Category.id
// 		WHERE Expense.userId = ?;
// 	`, userId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	expenses := make([]db.Expense, 0)

// 	for rows.Next() {
// 		var expense db.Expense
// 		var dateString string

// 		err := rows.Scan(
// 			&expense.Id,
// 			&dateString,
// 			&expense.Description,
// 			&expense.Amount,
// 			&expense.Notes,

// 			&expense.Currency.Id,
// 			&expense.Currency.Name,
// 			&expense.Currency.Symbol,

// 			&expense.Category.Id,
// 			&expense.Category.Name,
// 			&expense.Category.Description,

// 			&expense.SubCategory.Id,
// 			&expense.SubCategory.Name,
// 			&expense.SubCategory.Description,

// 			&expense.Purpose.Id,
// 			&expense.Purpose.Name,
// 			&expense.Purpose.Description,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		expense.Date, err = time.Parse("2006-01-02 03:04:05-07:00", dateString)
// 		if err != nil {
// 			// return nil, errors.New(fmt.Sprintf("Date format incorrect for %s", dateString))
// 			// if this happens, then something has been royally fucked up XD
// 			continue
// 		}

// 		// for each expense, we need to collate all its tags
// 		var tags []db.Tag
// 		etRows, err := d.DbConn().Query(`
// 			SELECT
// 				Tag.id,
// 				Tag.name,
// 				Tag.description
// 			FROM ExpenseTag
// 				JOIN Tag ON ExpenseTag.tagId = Tag.id
// 			WHERE userId = ? AND expenseId = ?
// 			,
// 		`, userId, expense.Id)

// 		for etRows.Next() {
// 			var tag db.Tag

// 			err = rows.Scan(&tag.Id, &tag.Name, &tag.Description)
// 		}

// 		expenses = append(expenses, expense)
// 	}

// 	return expenses, nil
// }
