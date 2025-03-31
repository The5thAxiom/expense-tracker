package server

import (
	"backend/db"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
)

func filterPaymentsService(w http.ResponseWriter, payments []db.Payment, queryParams url.Values) {
	filteredPayments := append(make([]db.Payment, len(payments)), payments...)
	appliedFilters := make([]string, 0)
	filters := []string{
		"amountIs", "amountMin", "amountMax",
		"dateIs", "dateBefore", "dateAfter",
		"descriptionIs", "descriptionIsNot", "descriptionIncludes", "descriptionDoesNoInclude", "descriptionLike",
		"categoryIs", "categoryIsNot",
		"subCategoryIs", "subCategoryIsNot",
		"purposeIs", "purposeIsNot",
		"currencyIs", "currencyIsNot",
		"notesIs", "notesIsNot", "notesIncludes", "notesDoesNotInclude", "notesLike",
	}
	for filter, values := range queryParams {
		if !slices.Contains(filters, filter) {
			continue
		}

		// amount filters
		if filter == "amountIs" {
			if len(values) > 1 {
				errorMessage := "Only one value allowed for 'amountIs'"
				NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}
			amountVal, err := strconv.ParseFloat(values[0], 64)
			if err != nil {
				errorMessage := fmt.Sprintf("Value for amountIs (%s) should be a float", values[0])
				NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			tempPayments := make([]db.Payment, 0)
			for _, payment := range filteredPayments {
				if payment.Amount == amountVal {
					tempPayments = append(tempPayments, payment)
				}
			}
			filteredPayments = tempPayments
			appliedFilters = append(appliedFilters, filter)
		}

		if filter == "amountMin" {
			if slices.Contains(appliedFilters, "amountIs") {
				errorMessage := "Cannot have both 'amountMin' and 'amountIs' filters together"
				NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			if len(values) > 1 {
				errorMessage := "Only one value allowed for 'amountMin'"
				NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			amountVal, err := strconv.ParseFloat(values[0], 64)
			if err != nil {
				errorMessage := fmt.Sprintf("Value for amountMin (%s) should be a float", values[0])
				NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			// if slices.Contains(appliedFilters, "amountMax") && queryParams.Get("amountMax") > amountVal {
			// 	errorMessage := "Value of 'amountMin' cannot be higher than 'amountMax'"
			// 	NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
			// 	return
			// }

			tempPayments := make([]db.Payment, 0)
			for _, payment := range filteredPayments {
				if payment.Amount >= amountVal {
					tempPayments = append(tempPayments, payment)
				}
			}
			filteredPayments = tempPayments
		}

		if filter == "amountMax" {
			if slices.Contains(appliedFilters, "amountIs") {
				errorMessage := "Cannot have both 'amountMax' and 'amountIs' filters together"
				NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			if len(values) > 1 {
				errorMessage := "Only one value allowed for 'amountMax'"
				NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			amountVal, err := strconv.ParseFloat(values[0], 64)
			if err != nil {
				errorMessage := fmt.Sprintf("Value for amountMax (%s) should be a float", values[0])
				NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			// if slices.Contains(appliedFilters, "amountMin") && queryParams.Get("amountMin") > amountVal {
			// 	amountMinVal
			// 	errorMessage := "Value of 'amountMin' cannot be higher than 'amountMax'"
			// 	NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
			// 	return
			// }

			tempPayments := make([]db.Payment, 0)
			for _, payment := range filteredPayments {
				if payment.Amount <= amountVal {
					tempPayments = append(tempPayments, payment)
				}
			}
			filteredPayments = tempPayments
		}

		if filter == "descriptionIncludes" {
			// if slices.Contains(appliedFilters, "amountIs") {
			// 	errorMessage := "Cannot have both 'amountMax' and 'amountIs' filters together"
			// 	NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
			// 	return
			// }

			if len(values) > 1 {
				errorMessage := "Only one value allowed for 'amountMax'"
				NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			value := values[0]

			tempPayments := make([]db.Payment, 0)
			for _, payment := range filteredPayments {
				if strings.Contains(strings.ToLower(payment.Description), strings.ToLower(value)) {
					tempPayments = append(tempPayments, payment)
				}
			}
			filteredPayments = tempPayments
		}

		// date filters

		// description filters
	}
	NewEndpointResponse(w, http.StatusOK, filteredPayments, nil)
}
