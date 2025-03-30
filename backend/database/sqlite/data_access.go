package sqlite

import (
	"backend/database"
	"database/sql"
)

func (d SQLiteDB) GetAllCategories() ([]database.Category, error) {
	db := d.Db()
	rows, err := db.Query(`SELECT id, name, description FROM Category;`)
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
	db := d.db

	var category database.Category
	var description sql.NullString

	err := db.QueryRow(`SELECT id, name, description FROM CATEGORY WHERE id=?;`, categoryId).Scan(&category.Id, &category.Name, &description)
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
