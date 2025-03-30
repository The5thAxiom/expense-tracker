package batch

import (
	"log"
	"strconv"
	"strings"
	"time"

	excelize "github.com/xuri/excelize/v2"
)

func ReadPayments(filename string, sheetname string) ([]ExcelPaymentRow, error) {
	rows, err := readRows(filename, sheetname)
	if err != nil {
		return nil, err
	}

	rows = rows[1:] // ignore the header

	payments := make([]ExcelPaymentRow, len(rows))

	var runningDate *time.Time = nil
	paymentIndexForRunningDate := 0

	for index, row := range rows {
		if len(row) < 7 {
			log.Printf("row #%d has empty values", index)
			break
		}

		dateString := row[1]
		timeObject, err := time.Parse("2-Jan-06", dateString)
		if err != nil {
			log.Fatalf("Could not parse time (%s) for row# %d: %s", dateString, index, err.Error())
		}

		if runningDate != nil && *runningDate == timeObject {
			paymentIndexForRunningDate += 1
		} else {
			runningDate = &timeObject
			paymentIndexForRunningDate = 0
		}

		amountString := row[4]
		amount, err := parseAmount(amountString)
		if err != nil {
			log.Fatalf("Could not parse amount (%s) for row# %d: %s", amountString, index, err.Error())
		}

		var purpose *string = nil
		if len(row) > 7 {
			purpose = &row[7]
		}

		var notes *string = nil
		if len(row) > 8 {
			notes = &row[8]
		}

		payments[index] = ExcelPaymentRow{
			Date:         timeObject,
			PaymentIndex: paymentIndexForRunningDate,
			Description:  row[3],
			Amount:       amount,
			Category:     row[5],
			Currency:     "INR",
			SubCategory:  row[6],
			Purpose:      purpose,
			Notes:        notes,
		}
	}

	return payments, nil
}

func readRows(filename string, sheetname string) ([][]string, error) {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	rows, err := file.GetRows(sheetname)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func parseAmount(amountString string) (float64, error) {
	amountString = strings.TrimSpace(amountString)
	amountString = strings.TrimPrefix(amountString, "₹")
	amountString = strings.TrimSpace(amountString)
	amountString = strings.Replace(amountString, ",", "", -1)

	amount := float64(0)
	var err error = nil

	if len(amountString) > 0 {
		amount, err = strconv.ParseFloat(amountString, 64)
	}

	return amount, err
}
