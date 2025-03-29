package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func Import(excelFilename string, sheetName string, db *sql.DB) (int, error) {
	payments, err := ReadPayments(excelFilename, sheetName)

	if err != nil {
		return 0, err
	}

	jsonBytes, _ := json.Marshal(payments[635:700])

	fmt.Println(string(jsonBytes))

	return len(payments), nil
}
