package backend

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

	db.Exec("PRAGMA foreign_keys = ON;")

	return db, nil
}
