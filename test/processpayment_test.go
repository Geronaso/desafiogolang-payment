// handlers_test.go
// Este arquivo contém testes para o handler ProcessPayment do serviço de pagamentos.
// Ele verifica se o processamento de pagamentos é realizado corretamente com base nas requisições recebidas.
// Utiliza a biblioteca testify/assert para validação dos resultados e net/http/httptest para simular requisições HTTP.

// O arquivo inclui três testes principais:
// 1. TestProcessPayment_ValidRequest: Verifica se uma solicitação válida resulta em uma resposta de sucesso com o ID da transação.
// 2. TestProcessPayment_InvalidRequest: Verifica se uma solicitação inválida (faltando dados) resulta em um erro adequado.
// 3. TestProcessPayment_UnsupportedGateway: Verifica se o uso de um gateway não suportado resulta em um erro adequado.

package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"desafiogolang-payment/handlers"
	"desafiogolang-payment/models"

	"github.com/stretchr/testify/assert"
)

func TestProcessPayment_ValidRequest(t *testing.T) {
	// Cria uma solicitação válida
	paymentRequest := models.PaymentRequest{
		Gateway:       "PayPal",
		Amount:        100.00,
		Currency:      "USD",
		PaymentMethod: "credit_card",
		CardDetails: models.CardDetails{
			Number: "4111111111111111",
			Expiry: "12/25",
			CVV:    "123",
		},
	}
	reqBody, _ := json.Marshal(paymentRequest)
	req, err := http.NewRequest("POST", "/process-payment", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ProcessPayment)

	// Chama o handler
	handler.ServeHTTP(rr, req)

	// Verifica o status da resposta
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verifica a resposta JSON
	var response models.PaymentResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Contains(t, response.Message, "Payment processed with success")
	assert.NotEmpty(t, response.Transaction_ID)
}

func TestProcessPayment_InvalidRequest(t *testing.T) {
	// Cria uma solicitação inválida
	invalidRequest := []byte(`{}`)
	req, err := http.NewRequest("POST", "/process-payment", bytes.NewBuffer(invalidRequest))
	if err != nil {
		t.Fatal(err)
	}

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ProcessPayment)

	// Chama o handler
	handler.ServeHTTP(rr, req)

	// Verifica o status da resposta
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Invalid request data\n", rr.Body.String())
}

func TestProcessPayment_UnsupportedGateway(t *testing.T) {
	// Cria uma solicitação com um gateway não suportado
	paymentRequest := models.PaymentRequest{
		Gateway:       "Stonego",
		Amount:        100.00,
		Currency:      "USD",
		PaymentMethod: "credit_card",
		CardDetails: models.CardDetails{
			Number: "4111111111111111",
			Expiry: "12/25",
			CVV:    "123",
		},
	}
	reqBody, _ := json.Marshal(paymentRequest)
	req, err := http.NewRequest("POST", "/process-payment", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ProcessPayment)

	// Chama o handler
	handler.ServeHTTP(rr, req)

	// Verifica o status da resposta
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Unsupported gateway\n", rr.Body.String())
}
