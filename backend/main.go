package main

import (
	"backend/batch"
	"backend/db/sqlite"
	"backend/server"
	"backend/vars"
	"fmt"
	"log"
	"os"
)

func main() {
	vars.Init()

	dbName := vars.Get[string]("db")
	resetDb := vars.Get[bool]("reset-db")

	db, err := sqlite.SQLite(dbName, resetDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.DbConn().Close()

	command := os.Args[1]
	switch command {
	case "init":
		fmt.Println("Initialized db")
	case "import-excel":
		excelFileName := vars.Get[string]("excel")
		sheetName := vars.Get[string]("sheet")

		num, err := batch.ImportExcelToDb(excelFileName, sheetName, db)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Added %d expenses to %s from %s", num, dbName, excelFileName)
	case "serve":
		server := server.New("0.0.0.0", 8000, db, nil)
		server.Run()
	case "help":
		printHelp()
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Expense Tracker Backend")
	fmt.Println("Valid commands :\n\tinit: To initialize the database and table\n\timport-excel: To import expenses from an excel sheet and write them to the database\n\tserve: To run the server\n\thelp: To display this help message")
	fmt.Println("Valid flags: \n\t--reset-db: To delete the database and create a new one")
	fmt.Println("\timport-excel:\n\t\t--excel <path-to-excel-sheet>: The path to the excel sheet to import expenses from (REQUIRED)\n\t\t--sheet <sheet-name>: The name of the sheet which has the expense data (REQUIRED)")
	fmt.Println("\tserve:\n\t\t--port <port-number>: The port on which to run the server (defaults to 8000)")
}
