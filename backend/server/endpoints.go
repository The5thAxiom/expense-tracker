package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s Server) AddEndpoints() {
	s.Mux.HandleFunc("/", s.index)

	s.Mux.HandleFunc("/categories", s.UseMiddleware(s.categories))
	s.Mux.HandleFunc("/categories/{categoryId}", s.UseMiddleware(s.categoryById))

	s.Mux.HandleFunc("/categories/{categoryId}/sub-categories", s.UseMiddleware(s.subCategories))
	s.Mux.HandleFunc("/categories/{categoryId}/sub-categories/{subCategoryId}", s.UseMiddleware(s.subCategoryById))

	s.Mux.HandleFunc("/purposes", s.UseMiddleware(s.purposes))
	s.Mux.HandleFunc("/purposes/{purposeId}", s.UseMiddleware(s.purposesById))

	s.Mux.HandleFunc("/currencies", s.UseMiddleware(s.currencies))
	s.Mux.HandleFunc("/currencies/{currencyAbbreviation}", s.UseMiddleware(s.currencyByAbbreviation))

	s.Mux.HandleFunc("/payments", s.UseMiddleware(s.payments))
	s.Mux.HandleFunc("/payments/{paymentId}", s.UseMiddleware(s.paymentById))
}

func (s Server) index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to the API")
}

func (s Server) categories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		categories, err := s.Db.GetAllCategories()
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		NewResponse(w, http.StatusOK, categories, nil)
	default:
		NewResponse(w, http.StatusBadRequest, nil, nil)
	}
}

func (s Server) categoryById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		categoryId := r.PathValue("categoryId")
		category, err := s.Db.GetCategoryById(categoryId)
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		if category == nil {
			errorMessage := "No category found for id '" + categoryId + "'"
			NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
			return
		}

		NewResponse(w, http.StatusOK, *category, nil)
	default:
		NewResponse(w, http.StatusBadRequest, nil, nil)
	}
}

func (s Server) subCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		categoryId := r.PathValue("categoryId")
		subCategories, err := s.Db.GetAllSubCategoriesforCategory(categoryId)
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		NewResponse(w, http.StatusOK, subCategories, nil)
	default:
		NewResponse(w, http.StatusBadRequest, nil, nil)
	}
}

func (s Server) subCategoryById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		categoryId := r.PathValue("categoryId")
		subCategoryId := r.PathValue("subCategoryId")

		category, err := s.Db.GetSubCategoryForCategoryById(subCategoryId, categoryId)
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		if category == nil {
			errorMessage := "No sub category found for id '" + categoryId + "'"
			NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
			return
		}

		NewResponse(w, http.StatusOK, *category, nil)
	default:
		NewResponse(w, http.StatusBadRequest, nil, nil)
	}
}

func (s Server) currencies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		currencies, err := s.Db.GetAllCurrencies()
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		NewResponse(w, http.StatusOK, currencies, nil)
	default:
		NewResponse(w, http.StatusBadRequest, nil, nil)
	}
}

func (s Server) currencyByAbbreviation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		currencyAbbreviation := r.PathValue("currencyAbbreviation")
		currency, err := s.Db.GetCurrencyByAbbreviation(currencyAbbreviation)
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		if currency == nil {
			errorMessage := "No currency found for id '" + currencyAbbreviation + "'"
			NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
			return
		}

		NewResponse(w, http.StatusOK, *currency, nil)
	default:
		NewResponse(w, http.StatusBadRequest, nil, nil)
	}
}

func (s Server) purposes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		purposes, err := s.Db.GetAllPurposes()
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		NewResponse(w, http.StatusOK, purposes, nil)
	default:
		NewResponse(w, http.StatusBadRequest, nil, nil)
	}
}

func (s Server) purposesById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		purposeId := r.PathValue("purposeId")
		purpose, err := s.Db.GetPurposeById(purposeId)
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		if purpose == nil {
			errorMessage := "No payment purpose found for id '" + purposeId + "'"
			NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
			return
		}

		NewResponse(w, http.StatusOK, *purpose, nil)
	default:
		NewResponse(w, http.StatusBadRequest, nil, nil)
	}
}

func (s Server) payments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		queryParams := r.URL.Query()

		payments, err := s.Db.GetAllPayments()
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		if len(queryParams) == 0 {
			NewResponse(w, http.StatusOK, payments, nil)
		} else {
			filterPaymentsService(w, payments, queryParams)
		}
	default:
		NewResponse(w, http.StatusBadRequest, nil, nil)
	}
}

func (s Server) paymentById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		paymentIdPathVar := r.PathValue("paymentId")
		paymentIdInt64, err := strconv.ParseInt(paymentIdPathVar, 10, 32)
		paymentId := int(paymentIdInt64)
		if err != nil {
			errorMessage := fmt.Sprintf("Payment Id '%s' should be an integer", paymentIdPathVar)
			NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
			return
		}

		payment, err := s.Db.GetPaymentById(paymentId)
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		if payment == nil {
			errorMessage := "No payment found for id '" + paymentIdPathVar + "'"
			NewResponse(w, http.StatusBadRequest, nil, &errorMessage)
			return
		}

		NewResponse(w, http.StatusOK, *payment, nil)
	default:
		NewResponse(w, http.StatusBadRequest, nil, nil)
	}
}
