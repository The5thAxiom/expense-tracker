package sqlite

import (
	"backend/db"
	"database/sql"
)

func (d SQLiteDB) GetCategoryById(userId string, id string) (*db.Category, error) {
	var category db.Category
	var description sql.NullString

	err := d.dbConn.QueryRow(
		`SELECT id, name, description FROM Category WHERE userId=? AND id=?;`, userId, id,
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
