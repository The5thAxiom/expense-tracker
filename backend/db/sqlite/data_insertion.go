package sqlite

import (
	"backend/db"
	"errors"
)

func (d SQLiteDB) InsertCategory(category db.CategoryEntity) error {
	_, err := d.dbConn.Exec(
		`INSERT INTO Category (userId, id, name, description) VALUES (?, ?, ?, ?);`,
		category.UserId, category.Id, category.Name, category.Description,
	)
	if err != nil {
		return errors.New("Error inserting into Category table: " + err.Error())
	}

	return nil
}

func (d SQLiteDB) InsertSubCategory(subCategory db.SubCategoryEntity) error {
	_, err := d.dbConn.Exec(
		`INSERT INTO SubCategory (userId, id, name, categoryId, description) VALUES (?, ?, ?, ?, ?);`,
		subCategory.UserId, subCategory.Id, subCategory.Name, subCategory.CategoryId, subCategory.Description,
	)
	if err != nil {
		return errors.New("Error inserting into SubCategory table: " + err.Error())
	}
	return nil
}

func (d SQLiteDB) InsertPurpose(newPurpose db.PurposeEntity) error {
	_, err := d.dbConn.Exec(
		`INSERT INTO Purpose (userId, id, name, description) VALUES (?, ?, ?, ?);`,
		newPurpose.UserId, newPurpose.Id, newPurpose.Name, newPurpose.Description,
	)
	if err != nil {
		return errors.New("Error inserting into Purpose table: " + err.Error())
	}

	return nil
}

func (d SQLiteDB) InsertCurrency(newCurrency db.CurrencyEntity) error {
	_, err := d.dbConn.Exec(
		`INSERT INTO Currency (id, name, symbol) VALUES (?, ?, ?);`,
		newCurrency.Id, newCurrency.Name, newCurrency.Symbol,
	)
	if err != nil {
		return errors.New("Error inserting into Currency table: " + err.Error())
	}

	return nil
}

func (d SQLiteDB) InsertTag(newTag db.TagEntity) error {
	_, err := d.dbConn.Exec(
		`INSERT INTO Tag (userId, id, name, description) VALUES (?, ?, ?, ?);`,
		newTag.UserId, newTag.Id, newTag.Name, newTag.Description,
	)
	if err != nil {
		return errors.New("Error inserting into Tag table: " + err.Error())
	}

	return nil
}

func (d SQLiteDB) LinkTagToExpense(tagId string, expenseId string) error {
	_, err := d.dbConn.Exec(
		`INSERT INTO ExpenseTag (tagId, expenseId) VALUES (?, ?);`,
		tagId, expenseId,
	)
	if err != nil {
		return errors.New("Error inserting into ExpenseTag table: " + err.Error())
	}

	return nil
}

func (d SQLiteDB) InsertExpense(userId string, expense db.ExpenseEntity) error {
	_, err := d.dbConn.Exec(`
		INSERT INTO Expense (
			userId,
			id,
			date,
			description,
			amount,
			currencyId,
			subCategoryId,
			purposeId,
			notes
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
	`,
		userId,
		expense.Id,
		expense.Date,
		expense.Description,
		expense.Amount,
		expense.CurrencyId,
		expense.SubCategoryId,
		expense.PurposeId,
		expense.Notes,
	)
	if err != nil {
		return err
	}

	return nil
}
