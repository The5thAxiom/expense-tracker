package sqlite

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	dbConn     *sql.DB
	dbFileName string
}

func (d SQLiteDB) DbConn() *sql.DB {
	return d.dbConn
}

func SQLite(dbFileName string, resetDb bool) (SQLiteDB, error) {
	var sqliteDb SQLiteDB
	sqliteDb.dbFileName = dbFileName

	if resetDb {
		log.Println("Deleting database file")
		err := os.Remove(dbFileName)
		if err != nil {
			return sqliteDb, err
		}
	}

	sqliteDb.createDbIfNotExist()

	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		return sqliteDb, err
	}
	sqliteDb.dbConn = db

	err = db.Ping()
	if err != nil {
		return sqliteDb, errors.New("Could not ping database " + err.Error())
	}

	db.Exec("PRAGMA foreign_keys = ON;")

	err = sqliteDb.createTablesIfNotExist()
	if err != nil {
		return sqliteDb, errors.New("Could not create tables: " + err.Error())
	}

	return sqliteDb, nil
}

func (d SQLiteDB) createDbIfNotExist() {
	_, err := os.Stat(d.dbFileName)
	if err != nil {
		log.Printf("Database file '%s' does not exist.\nCreating...", d.dbFileName)
		file, fileCreationErr := os.Create(d.dbFileName)
		if fileCreationErr != nil {
			log.Fatalf("Could not create database file '%s'", d.dbFileName)
		}
		defer file.Close()
		log.Printf("Created file '%s'", d.dbFileName)
	}
}

func (d SQLiteDB) createTablesIfNotExist() error {
	db := d.dbConn
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
