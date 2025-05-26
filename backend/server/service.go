package server

import (
	"backend/db"
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func newSqlNullString(str *string) sql.NullString {
	if str != nil {
		return sql.NullString{
			String: *str,
			Valid:  true,
		}
	} else {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}
}

func (s Server) addExpense(userId string, newExpense NewExpense) (string, error) {
	var categoryId string
	var subCategoryId string
	var purposeId sql.NullString
	var tagIds []string = make([]string, len(newExpense.Tags))

	// insert category if new
	if newExpense.Category.IsNew {
		categoryId = uuid.New().String()
		newCategory := db.CategoryEntity{
			Id:          categoryId,
			UserId:      userId,
			Name:        *newExpense.Category.Name,
			Description: newSqlNullString(newExpense.Category.Description),
		}
		err := s.Db.InsertCategory(newCategory)
		if err != nil {
			return "", err
		}
	} else {
		categoryId = *newExpense.Category.Id
		// category, err := s.Db.GetCategoryById(userId, newExpense.Currency.Id)
		// if err != nil {
		// 	return "", errors.New("could not find category ")
		// }
	}

	// insert sub category if new
	if newExpense.SubCategory.IsNew {
		subCategoryId = uuid.New().String()
		newSubCategory := db.SubCategoryEntity{
			Id:          subCategoryId,
			UserId:      userId,
			CategoryId:  categoryId,
			Name:        *newExpense.SubCategory.Name,
			Description: newSqlNullString(newExpense.SubCategory.Description),
		}
		err := s.Db.InsertSubCategory(newSubCategory)
		if err != nil {
			return "", err
		}
	} else {
		subCategoryId = *newExpense.SubCategory.Id
	}

	// insert tags if new
	for i, tag := range newExpense.Tags {
		if tag.IsNew {
			tagId := uuid.New().String()
			newTag := db.TagEntity{
				Id:          tagId,
				UserId:      userId,
				Name:        *tag.Name,
				Description: newSqlNullString(tag.Description),
			}
			err := s.Db.InsertTag(newTag)
			if err != nil {
				return "", err
			}
			tagIds[i] = newTag.Id
		} else {
			tagIds[i] = *tag.Id
		}
	}

	// insert purpose if new
	if newExpense.Purpose != nil {
		if newExpense.Purpose.IsNew {
			purposeId := uuid.New().String()
			newPurpose := db.PurposeEntity{
				Id:          purposeId,
				UserId:      userId,
				Name:        *newExpense.Purpose.Name,
				Description: newSqlNullString(newExpense.Purpose.Description),
			}
			err := s.Db.InsertPurpose(newPurpose)
			if err != nil {
				return "", err
			}
		} else {
			purposeId = newSqlNullString(newExpense.Purpose.Id)
		}
	}

	// insert expense into db
	expenseId := uuid.New().String()
	expense := db.ExpenseEntity{
		Id:            expenseId,
		Date:          newExpense.Date,
		Description:   newExpense.Description,
		Amount:        newExpense.Amount,
		Notes:         newSqlNullString(newExpense.Notes),
		CurrencyId:    newExpense.CurrencyId,
		SubCategoryId: subCategoryId,
		PurposeId:     purposeId,
	}
	err := s.Db.InsertExpense(userId, expense)
	if err != nil {
		return "", err
	}

	// link tags to the expense
	for _, tagId := range tagIds {
		s.Db.LinkTagToExpense(tagId, expenseId)
	}

	return expenseId, nil
}

func filterExpensesService(w http.ResponseWriter, expenses []db.Expense, queryParams url.Values) {
	filteredExpenses := append(make([]db.Expense, len(expenses)), expenses...)
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
				NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}
			amountVal, err := strconv.ParseInt(values[0], 10, 64)
			if err != nil {
				errorMessage := fmt.Sprintf("Value for amountIs (%s) should be a float", values[0])
				NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			tempExpenses := make([]db.Expense, 0)
			for _, expense := range filteredExpenses {
				if expense.Amount == int(amountVal) {
					tempExpenses = append(tempExpenses, expense)
				}
			}
			filteredExpenses = tempExpenses
			appliedFilters = append(appliedFilters, filter)
		}

		if filter == "amountMin" {
			if slices.Contains(appliedFilters, "amountIs") {
				errorMessage := "Cannot have both 'amountMin' and 'amountIs' filters together"
				NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			if len(values) > 1 {
				errorMessage := "Only one value allowed for 'amountMin'"
				NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			amountVal, err := strconv.ParseInt(values[0], 10, 64)
			if err != nil {
				errorMessage := fmt.Sprintf("Value for amountMin (%s) should be a float", values[0])
				NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			// if slices.Contains(appliedFilters, "amountMax") && queryParams.Get("amountMax") > amountVal {
			// 	errorMessage := "Value of 'amountMin' cannot be higher than 'amountMax'"
			// 	NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
			// 	return
			// }

			tempExpenses := make([]db.Expense, 0)
			for _, expense := range filteredExpenses {
				if expense.Amount >= int(amountVal) {
					tempExpenses = append(tempExpenses, expense)
				}
			}
			filteredExpenses = tempExpenses
		}

		if filter == "amountMax" {
			if slices.Contains(appliedFilters, "amountIs") {
				errorMessage := "Cannot have both 'amountMax' and 'amountIs' filters together"
				NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			if len(values) > 1 {
				errorMessage := "Only one value allowed for 'amountMax'"
				NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			amountVal, err := strconv.ParseInt(values[0], 10, 64)
			if err != nil {
				errorMessage := fmt.Sprintf("Value for amountMax (%s) should be a float", values[0])
				NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			// if slices.Contains(appliedFilters, "amountMin") && queryParams.Get("amountMin") > amountVal {
			// 	amountMinVal
			// 	errorMessage := "Value of 'amountMin' cannot be higher than 'amountMax'"
			// 	NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
			// 	return
			// }

			tempExpenses := make([]db.Expense, 0)
			for _, expense := range filteredExpenses {
				if expense.Amount <= int(amountVal) {
					tempExpenses = append(tempExpenses, expense)
				}
			}
			filteredExpenses = tempExpenses
		}

		if filter == "descriptionIncludes" {
			// if slices.Contains(appliedFilters, "amountIs") {
			// 	errorMessage := "Cannot have both 'amountMax' and 'amountIs' filters together"
			// 	NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
			// 	return
			// }

			if len(values) > 1 {
				errorMessage := "Only one value allowed for 'amountMax'"
				NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
				return
			}

			value := values[0]

			tempExpenses := make([]db.Expense, 0)
			for _, expense := range filteredExpenses {
				if strings.Contains(strings.ToLower(expense.Description), strings.ToLower(value)) {
					tempExpenses = append(tempExpenses, expense)
				}
			}
			filteredExpenses = tempExpenses
		}

		// date filters

		// description filters
	}
	NewResponse(w, http.StatusOK, filteredExpenses, nil)
}
