// currency_conversion_test.go
// Este arquivo contém testes para o handler ConvertCurrency do serviço de conversão de moeda.
// Ele verifica se a conversão de moeda é realizada corretamente com base nas requisições recebidas.
// Utiliza a biblioteca testify/assert para validação dos resultados e net/http/httptest para simular requisições HTTP.

// O arquivo inclui três testes principais:
// 1. TestConvertCurrency_ValidRequest: Verifica se uma solicitação válida resulta em uma resposta de sucesso com o valor convertido.
// 2. TestConvertCurrency_InvalidRequest: Verifica se uma solicitação inválida (faltando dados) resulta em um erro adequado.
// 3. TestConvertCurrency_ExternalAPIError: Verifica se um erro na API externa resulta em um erro adequado.

package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"desafiogolang-payment/handlers"
	"desafiogolang-payment/models"
	"desafiogolang-payment/services"

	"github.com/stretchr/testify/assert"
)

// Mock service response for ConvertCurrency
func mockConvertCurrency(request models.CurrencyConversionRequest) (models.CurrencyConversionResponse, error) {
	if request.FromCurrency == "USD" && request.ToCurrency == "EUR" {
		return models.CurrencyConversionResponse{
			ConvertedAmount: 85.00,
			FromCurrency:    "USD",
			ToCurrency:      "EUR",
			Rate:            0.85,
		}, nil
	}
	return models.CurrencyConversionResponse{}, fmt.Errorf("conversion error")
}

func TestConvertCurrency_ValidRequest(t *testing.T) {
	// Injeção de dependência da função mock
	originalFunc := services.ConvertCurrencyFunc
	services.ConvertCurrencyFunc = mockConvertCurrency
	defer func() { services.ConvertCurrencyFunc = originalFunc }()

	// Cria uma solicitação válida
	conversionRequest := models.CurrencyConversionRequest{
		Amount:       100.00,
		FromCurrency: "USD",
		ToCurrency:   "EUR",
	}
	reqBody, _ := json.Marshal(conversionRequest)
	req, err := http.NewRequest("POST", "/convert-currency", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ConvertCurrency)

	// Chama o handler
	handler.ServeHTTP(rr, req)

	// Verifica o status da resposta
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verifica a resposta JSON
	var response models.CurrencyConversionResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 85.00, response.ConvertedAmount)
	assert.Equal(t, "USD", response.FromCurrency)
	assert.Equal(t, "EUR", response.ToCurrency)
	assert.Equal(t, 0.85, response.Rate)
}

func TestConvertCurrency_InvalidRequest(t *testing.T) {
	// Cria uma solicitação inválida
	invalidRequest := []byte(`{}`)
	req, err := http.NewRequest("POST", "/convert-currency", bytes.NewBuffer(invalidRequest))
	if err != nil {
		t.Fatal(err)
	}

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ConvertCurrency)

	// Chama o handler
	handler.ServeHTTP(rr, req)

	// Verifica o status da resposta
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Invalid request data\n", rr.Body.String())
}

func TestConvertCurrency_ExternalAPIError(t *testing.T) {
	// Injeção de dependência da função mock
	originalFunc := services.ConvertCurrencyFunc
	services.ConvertCurrencyFunc = mockConvertCurrency
	defer func() { services.ConvertCurrencyFunc = originalFunc }()

	// Cria uma solicitação que resultará em um erro de conversão
	conversionRequest := models.CurrencyConversionRequest{
		Amount:       100.00,
		FromCurrency: "USD",
		ToCurrency:   "XXX",
	}
	reqBody, _ := json.Marshal(conversionRequest)
	req, err := http.NewRequest("POST", "/convert-currency", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ConvertCurrency)

	// Chama o handler
	handler.ServeHTTP(rr, req)

	// Verifica o status da resposta
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
