package main

import (
	batch "backend/batch"
	d "backend/db"
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

	db, err := d.SQLite(dbName, resetDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch command {
	case "export":
		excelFileName := getCliArgument("excel")
		sheetName := getCliArgument("sheet")

		num, err := batch.Import(excelFileName, sheetName, db)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Added %d payments to %s from %s", num, dbName, excelFileName)
	case "serve":
		fmt.Print("serving...")
	}
}
