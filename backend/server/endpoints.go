package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s Server) AddEndpoints() {
	s.Mux.HandleFunc("/", s.index)

	s.Mux.HandleFunc("/categories", s.categories)
	s.Mux.HandleFunc("/categories/{categoryId}", s.categoryById)

	s.Mux.HandleFunc("/sub-categories", s.subCategories)
	s.Mux.HandleFunc("/sub-categories/{subCategoryId}", s.subCategoryById)

	s.Mux.HandleFunc("/purposes", s.purposes)
	s.Mux.HandleFunc("/purposes/{purposeId}", s.purposesById)

	s.Mux.HandleFunc("/currencies", s.currencies)
	s.Mux.HandleFunc("/currencies/{currencyAbbreviation}", s.currencyByAbbreviation)

	s.Mux.HandleFunc("/payments", s.payments)
	s.Mux.HandleFunc("/payments/{paymentId}", s.paymentById)
}

func (s Server) index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to the API")
}

func (s Server) categories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		categories, err := s.Db.GetAllCategories()
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewEndpointResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		NewEndpointResponse(w, http.StatusOK, categories, nil)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s Server) categoryById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		categoryId := r.PathValue("categoryId")
		category, err := s.Db.GetCategoryById(categoryId)
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewEndpointResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		if category == nil {
			errorMessage := "No category found for id '" + categoryId + "'"
			NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
			return
		}

		NewEndpointResponse(w, http.StatusOK, *category, nil)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s Server) subCategories(w http.ResponseWriter, r *http.Request) {}

func (s Server) subCategoryById(w http.ResponseWriter, r *http.Request) {}

func (s Server) currencies(w http.ResponseWriter, r *http.Request) {}

func (s Server) currencyByAbbreviation(w http.ResponseWriter, r *http.Request) {}

func (s Server) purposes(w http.ResponseWriter, r *http.Request) {}

func (s Server) purposesById(w http.ResponseWriter, r *http.Request) {}

func (s Server) payments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		payments, err := s.Db.GetAllPayments()
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewEndpointResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		NewEndpointResponse(w, http.StatusOK, payments, nil)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s Server) paymentById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		paymentIdPathVar := r.PathValue("paymentId")
		paymentIdInt64, err := strconv.ParseInt(paymentIdPathVar, 10, 32)
		paymentId := int(paymentIdInt64)
		if err != nil {
			errorMessage := fmt.Sprintf("Payment Id '%s' should be an integer", paymentIdPathVar)
			NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
			return
		}

		payment, err := s.Db.GetPaymentById(paymentId)
		if err != nil {
			errorMessage := "An error occurred: " + err.Error()
			NewEndpointResponse(w, http.StatusInternalServerError, nil, &errorMessage)
			return
		}

		if payment == nil {
			errorMessage := "No payment found for id '" + paymentIdPathVar + "'"
			NewEndpointResponse(w, http.StatusBadRequest, nil, &errorMessage)
			return
		}

		NewEndpointResponse(w, http.StatusOK, *payment, nil)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
