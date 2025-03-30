package main

import (
	"backend/batch"
	"backend/db"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func getCliArgument(argument string) string {
	args := os.Args[1:]

	argumentNameIndex := slices.Index(args, "--"+argument)
	if argumentNameIndex == -1 {
		log.Fatalf("CLI argument not found, please provide '--%s <%s>'", argument, argument)
	}

	argumentIndex := argumentNameIndex + 1

	if len(args) < argumentIndex+1 || strings.HasPrefix(args[argumentIndex], "--") {
		log.Fatalf("No argument provided for flag '--%s'", argument)
	}

	return args[argumentIndex]
}

func getCliFlag(flag string) bool {
	return slices.Contains(os.Args[1:], "--"+flag)
}

func main() {
	command := os.Args[1]

	dbName := getCliArgument("db")
	resetDb := getCliFlag("reset-db")

	db, err := db.SQLite(dbName, resetDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch command {
	case "init":
		fmt.Println("Initialized db")
	case "import-excel":
		excelFileName := getCliArgument("excel")
		sheetName := getCliArgument("sheet")

		num, err := batch.ImportExcelToDb(excelFileName, sheetName, db)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Added %d payments to %s from %s", num, dbName, excelFileName)
	case "serve":
		fmt.Print("serving...")
	case "help":
		printHelp()
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Expense Tracker Backend")
	fmt.Println("Valid commands :\n\tinit: To initialize the database and table\n\timport-excel: To import payments from an excel sheet and write them to the database\n\tserve: To run the server\n\thelp: To display this help message")
	fmt.Println("Valid flags: \n\t--reset-db: To delete the database and create a new one")
	fmt.Println("\timport-excel:\n\t\t--excel <path-to-excel-sheet>: The path to the excel sheet to import payments from (REQUIRED)\n\t\t--sheet <sheet-name>: The name of the sheet which has the payment data (REQUIRED)")
	fmt.Println("\tserve:\n\t\t--port <port-number>: The port on which to run the server (defaults to 8000)")
}
