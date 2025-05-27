package batch

import (
	"backend/db"
	"fmt"
	"log"
)

func ImportExcelToDb(excelFilename string, sheetName string, d db.DB) (int, error) {
	expenses, err := ReadExpenses(excelFilename, sheetName)
	if err != nil {
		return 0, err
	}

	log.Printf("Read %d rows from %s[\"%s\"]", len(expenses), excelFilename, sheetName)

	failedWrites := make([]ExcelExpenseRow, 0)

	for i, p := range expenses {
		err = WriteExpense(d, p, i)
		if err != nil {
			log.Printf("Could not write expense #%d: %s", i, err.Error())
			failedWrites = append(failedWrites, p)
		}
	}

	if len(failedWrites) > 0 {
		log.Printf("Could not write %d expenses:", len(failedWrites))
		for i, p := range failedWrites {
			fmt.Printf("%d: %s", i, p.ToString())
		}
	}

	return len(expenses), nil
}
