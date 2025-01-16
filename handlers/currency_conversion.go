package handlers

import (
	"desafiogolang-payment/models"
	"desafiogolang-payment/services"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// ConvertCurrency lida com solicitações de conversão de moeda.
func ConvertCurrency(w http.ResponseWriter, r *http.Request) {
	var conversionRequest models.CurrencyConversionRequest

	if err := json.NewDecoder(r.Body).Decode(&conversionRequest); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(conversionRequest); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	response, err := services.ConvertCurrency(conversionRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}
