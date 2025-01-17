// currency_conversion.go
// Este arquivo contém o handler para processar solicitações de conversão de moeda.
// Ele utiliza os pacotes services e models para realizar a conversão de moeda e validar dados de entrada.
// As solicitações são recebidas como JSON, validadas e encaminhadas para o serviço apropriado para conversão.

// O arquivo inclui uma função principal:
// ConvertCurrency: Lida com solicitações de conversão de moeda, decodifica a solicitação JSON, valida os dados e encaminha para o serviço de conversão de moeda.

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
