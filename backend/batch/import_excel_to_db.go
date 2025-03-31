package batch

import (
	"backend/database"
	"fmt"
	"log"
)

func ImportExcelToDb(excelFilename string, sheetName string, d database.DB) (int, error) {
	payments, err := ReadPayments(excelFilename, sheetName)
	if err != nil {
		return 0, err
	}

	log.Printf("Read %d rows from %s[\"%s\"]", len(payments), excelFilename, sheetName)

	failedWrites := make([]ExcelPaymentRow, 0)

	for i, p := range payments {
		err = WritePayment(d, p, i)
		if err != nil {
			log.Printf("Could not write payment #%d: %s", i, err.Error())
			failedWrites = append(failedWrites, p)
		}
	}

	if len(failedWrites) > 0 {
		log.Printf("Could not write %d payments:", len(failedWrites))
		for i, p := range failedWrites {
			fmt.Printf("%d: %s", i, p.ToString())
		}
	}

	return len(payments), nil
}
