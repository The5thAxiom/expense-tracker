package server

import (
	"encoding/json"
	"fmt"

	// "fmt"

	"net/http"
	// "strconv"
)

func (s Server) AddEndpoints() {
	s.useMiddleware(s.logRequest)
	s.useMiddleware(s.loggedInUser)

	// s.addRoute("/api/test", s.testFunc)
	s.addRoute("/api/users", s.usersController)

	s.addRoute("/api/expenses", s.expensesController)
}

// func (s Server) testFunc(w http.ResponseWriter, r *http.Request) {
// 	userId := r.Header.Get("X-USER-ID")
// 	notes := "abcdef"
// 	catName := "dummyCat1"
// 	subCatName := "dummy subcat 1"

// 	newExpense := NewExpense{
// 		Date:        time.Now(),
// 		Description: "dummy payment added",
// 		Amount:      10_000,
// 		Notes:       &notes,
// 		CurrencyId:  "INR",
// 		Category: MaybeNew{
// 			IsNew: true,
// 			Name:  &catName,
// 		},
// 		SubCategory: MaybeNew{
// 			IsNew: true,
// 			Name:  &subCatName,
// 		},
// 	}

// 	id, err := s.addExpense(userId, newExpense)
// 	if err != nil {
// 		ErrorResponse(w, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	JsonResponse(w, http.StatusOK, Id{Id: id})
// }

func (s Server) expensesController(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("X-USER-ID")

	switch r.Method {
	case http.MethodGet:
		// 	queryParams := r.URL.Query()

		// 	expenses, err := s.Db.GetAllExpenses()
		// 	if err != nil {
		// 		errorMessage := "An error occurred: " + err.Error()
		// 		NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
		// 		return
		// 	}

		// 	if len(queryParams) == 0 {
		// 		NewResponse(w, http.StatusOK, expenses, nil)
		// 	} else {
		// 		filterExpensesService(w, expenses, queryParams)
		// 	}

		expenses, err := s.Db.GetAllExpenses(userId)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		JsonResponse(w, http.StatusOK, expenses)

	case http.MethodPost:
		var newExpense NewExpense

		err := json.NewDecoder(r.Body).Decode(&newExpense)
		if err != nil {
			ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		fmt.Printf("%s %T\n", newExpense.Date, newExpense.Date)

		id, err := s.addExpense(userId, newExpense)
		if err != nil {
			ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		JsonResponse(w, http.StatusCreated, Id{Id: id})
	default:
		NewResponse(w, http.StatusBadRequest, nil, nil)
	}
}

// func (s Server) expenseById(w http.ResponseWriter, r *http.Request) {
// 	userId := r.Header.Get("X-USER-ID")
// 	switch r.Method {
// 	case http.MethodGet:
// 		// expenseIdPathVar := r.PathValue("expenseId")
// 		// expenseIdInt64, err := strconv.ParseInt(expenseIdPathVar, 10, 32)
// 		// expenseId := int(expenseIdInt64)
// 		// if err != nil {
// 		// 	errorMessage := fmt.Sprintf("Expense Id '%s' should be an integer", expenseIdPathVar)
// 		// 	NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
// 		// 	return
// 		// }

// 		// expense, err := s.Db.GetExpenseById(expenseId)
// 		// if err != nil {
// 		// 	errorMessage := "An error occurred: " + err.Error()
// 		// 	NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
// 		// 	return
// 		// }

// 		// if expense == nil {
// 		// 	errorMessage := "No expense found for id '" + expenseIdPathVar + "'"
// 		// 	NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
// 		// 	return
// 		// }

// 		// NewResponse(w, http.StatusOK, *expense, nil)
// 	case http.MethodPost:

// 	default:
// 		NewResponse(w, http.StatusBadRequest, nil, nil)
// 	}
// }

func (s Server) usersController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		JsonResponse(w, http.StatusOK, s)
		return
	case http.MethodPut:
		JsonResponse(w, http.StatusOK, s)
		return
	case http.MethodPatch:
		JsonResponse(w, http.StatusOK, s)
		return
	case http.MethodPost:
		JsonResponse(w, http.StatusCreated, s)
		return
	case http.MethodDelete:
		JsonResponse(w, http.StatusOK, s)
		return
	default:
		ErrorResponse(w, http.StatusMethodNotAllowed, "Ye kaunsa method hai yar")
	}
}

// func (s Server) xAddEndpoints() {
// 	s.Mux.HandleFunc("/", s.index)

// 	s.Mux.HandleFunc("/categories", s.UseMiddleware(s.categories))
// 	s.Mux.HandleFunc("/categories/{categoryId}", s.UseMiddleware(s.categoryById))

// 	s.Mux.HandleFunc("/categories/{categoryId}/sub-categories", s.UseMiddleware(s.subCategories))
// 	s.Mux.HandleFunc("/categories/{categoryId}/sub-categories/{subCategoryId}", s.UseMiddleware(s.subCategoryById))

// 	s.Mux.HandleFunc("/purposes", s.UseMiddleware(s.purposes))
// 	s.Mux.HandleFunc("/purposes/{purposeId}", s.UseMiddleware(s.purposesById))

// 	s.Mux.HandleFunc("/currencies", s.UseMiddleware(s.currencies))
// 	s.Mux.HandleFunc("/currencies/{currencyId}", s.UseMiddleware(s.currencyById))

// 	s.Mux.HandleFunc("/expenses", s.UseMiddleware(s.expenses))
// 	s.Mux.HandleFunc("/expenses/{expenseId}", s.UseMiddleware(s.expenseById))
// }

// func (s Server) index(res http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(res, "Welcome to the API")
// }

// func (s Server) categories(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		categories, err := s.Db.GetAllCategories()
// 		if err != nil {
// 			ErrorResponse(w, http.StatusInternalServerError, err.Error())
// 			return
// 		}
// 		JsonResponse(w, http.StatusOK, categories)

// 	case http.MethodPost:
// 		var newCategory db.Category

// 		err := json.NewDecoder(r.Body).Decode(&newCategory)
// 		if err != nil {
// 			ErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		id, err := s.Db.AddCategory(newCategory)
// 		if err != nil {
// 			ErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		log.Printf("Added new category %s", id)

// 		data := struct {
// 			Id string `json:"id"`
// 		}{Id: id}
// 		JsonResponse(w, http.StatusCreated, data)

// 	default:
// 		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
// 	}
// }

// func (s Server) categoryById(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		categoryId := r.PathValue("categoryId")
// 		category, err := s.Db.GetCategoryById(categoryId)
// 		if err != nil {
// 			errorMessage := "An error occurred: " + err.Error()
// 			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
// 			return
// 		}

// 		if category == nil {
// 			errorMessage := "No category found for id '" + categoryId + "'"
// 			NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
// 			return
// 		}

// 		NewResponse(w, http.StatusOK, *category, nil)
// 	default:
// 		NewResponse(w, http.StatusBadRequest, nil, nil)
// 	}
// }

// func (s Server) subCategories(w http.ResponseWriter, r *http.Request) {
// 	categoryId := r.PathValue("categoryId")

// 	switch r.Method {
// 	case http.MethodGet:
// 		subCategories, err := s.Db.GetAllSubCategoriesforCategory(categoryId)
// 		if err != nil {
// 			errorMessage := "An error occurred: " + err.Error()
// 			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
// 			return
// 		}
// 		NewResponse(w, http.StatusOK, subCategories, nil)

// 	case http.MethodPost:
// 		var newSubCategory db.SubCategory

// 		err := json.NewDecoder(r.Body).Decode(&newSubCategory)
// 		if err != nil {
// 			ErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		id, err := s.Db.AddSubCategory(categoryId, newSubCategory)
// 		if err != nil {
// 			ErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		log.Printf("Added new sub category %s", id)

// 		data := struct {
// 			Id string `json:"id"`
// 		}{Id: id}
// 		JsonResponse(w, http.StatusCreated, data)

// 	default:
// 		NewResponse(w, http.StatusBadRequest, nil, nil)
// 	}
// }

// func (s Server) subCategoryById(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		categoryId := r.PathValue("categoryId")
// 		subCategoryId := r.PathValue("subCategoryId")

// 		category, err := s.Db.GetSubCategoryForCategoryById(subCategoryId, categoryId)
// 		if err != nil {
// 			errorMessage := "An error occurred: " + err.Error()
// 			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
// 			return
// 		}

// 		if category == nil {
// 			errorMessage := "No sub category found for id '" + categoryId + "'"
// 			NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
// 			return
// 		}

// 		NewResponse(w, http.StatusOK, *category, nil)
// 	default:
// 		NewResponse(w, http.StatusBadRequest, nil, nil)
// 	}
// }

// func (s Server) currencies(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		currencies, err := s.Db.GetAllCurrencies()
// 		if err != nil {
// 			errorMessage := "An error occurred: " + err.Error()
// 			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
// 			return
// 		}

// 		NewResponse(w, http.StatusOK, currencies, nil)

// 	case http.MethodPost:
// 		var newCategory db.Currency

// 		err := json.NewDecoder(r.Body).Decode(&newCategory)
// 		if err != nil {
// 			ErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		id, err := s.Db.AddCurrency(newCategory)
// 		if err != nil {
// 			ErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		log.Printf("Added new currency %s", id)

// 		data := struct {
// 			Id string `json:"id"`
// 		}{Id: id}
// 		JsonResponse(w, http.StatusCreated, data)
// 	default:
// 		NewResponse(w, http.StatusBadRequest, nil, nil)
// 	}
// }

// func (s Server) currencyById(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		currencyId := r.PathValue("currencyId")
// 		currency, err := s.Db.GetCurrencyById(currencyId)
// 		if err != nil {
// 			errorMessage := "An error occurred: " + err.Error()
// 			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
// 			return
// 		}

// 		if currency == nil {
// 			errorMessage := "No currency found for id '" + currencyId + "'"
// 			NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
// 			return
// 		}

// 		NewResponse(w, http.StatusOK, *currency, nil)
// 	default:
// 		NewResponse(w, http.StatusBadRequest, nil, nil)
// 	}
// }

// func (s Server) purposes(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		purposes, err := s.Db.GetAllPurposes()
// 		if err != nil {
// 			errorMessage := "An error occurred: " + err.Error()
// 			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
// 			return
// 		}

// 		NewResponse(w, http.StatusOK, purposes, nil)

// 	case http.MethodPost:
// 		var newPurpose db.Purpose

// 		err := json.NewDecoder(r.Body).Decode(&newPurpose)
// 		if err != nil {
// 			ErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		id, err := s.Db.AddPurpose(newPurpose)
// 		if err != nil {
// 			ErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		log.Printf("Added new purpose %s", id)

// 		data := struct {
// 			Id string `json:"id"`
// 		}{Id: id}
// 		JsonResponse(w, http.StatusCreated, data)
// 	default:
// 		NewResponse(w, http.StatusBadRequest, nil, nil)
// 	}
// }

// func (s Server) purposesById(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		purposeId := r.PathValue("purposeId")
// 		purpose, err := s.Db.GetPurposeById(purposeId)
// 		if err != nil {
// 			errorMessage := "An error occurred: " + err.Error()
// 			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
// 			return
// 		}

// 		if purpose == nil {
// 			errorMessage := "No expense purpose found for id '" + purposeId + "'"
// 			NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
// 			return
// 		}

// 		NewResponse(w, http.StatusOK, *purpose, nil)
// 	default:
// 		NewResponse(w, http.StatusBadRequest, nil, nil)
// 	}
// }
