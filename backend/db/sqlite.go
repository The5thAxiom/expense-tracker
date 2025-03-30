package db

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func createDbIfNotExist(dbFileName string) {
	_, err := os.Stat(dbFileName)
	if err != nil {
		log.Printf("Database file '%s' does not exist.\nCreating...", dbFileName)
		file, fileCreationErr := os.Create(dbFileName)
		if fileCreationErr != nil {
			log.Fatalf("Could not create database file '%s'", dbFileName)
		}
		defer file.Close()
		log.Printf("Created file '%s'", dbFileName)
	}
}

func createTablesIfNotExist(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Category (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT
		);
	`)
	if err != nil {
		return errors.New("Category: " + err.Error())
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS SubCategory (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			categoryId TEXT NOT NULL,
			FOREIGN KEY (categoryId) REFERENCES Category(id)
		);
	`)
	if err != nil {
		return errors.New("SubCategory: " + err.Error())
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Purpose (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT
		);
	`)
	if err != nil {
		return errors.New("Purpose: " + err.Error())
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Currency (
			abbreviation TEXT PRIMARY KEY,
			name TEXT,
			symbol TEXT
		);
	`)
	if err != nil {
		return errors.New("Currency: " + err.Error())
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Payment (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date TEXT NOT NULL,
			paymentIndex INTEGER NOT NULL,
			description TEXT NOT NULL,
			amount INTEGER NOT NULL,
			currencyAbbreviation INTEGER NOT NULL,
			subCategoryId INTEGER NOT NULL,
			purposeId INTEGER,
			notes TEXT,
			FOREIGN KEY (currencyAbbreviation) REFERENCES Currency(abbreviation),
			FOREIGN KEY (subCategoryId) REFERENCES SubCategory(id),
			FOREIGN KEY (purposeId) REFERENCES Purpose(id)
		);
	`)
	if err != nil {
		return errors.New("FullCategory: " + err.Error())
	}

	return nil
}

func SQLite(dbFileName string, resetDb bool) (*sql.DB, error) {
	if resetDb {
		os.Remove(dbFileName)
	}

	createDbIfNotExist(dbFileName)

	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("Could not ping database " + err.Error())
	}

	err = createTablesIfNotExist(db)
	if err != nil {
		return nil, errors.New("Could not create tables: " + err.Error())
	}

	db.Exec("PRAGMA foreign_keys = ON;")

	return db, nil
}
