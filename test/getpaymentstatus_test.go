// handlers_test.go
// Este arquivo contém testes para o handler GetPaymentStatus do serviço de pagamentos.
// Ele verifica se o status de uma transação é corretamente retornado com base no ID da transação e no gateway de pagamento fornecido.
// Utiliza a biblioteca testify/assert para validação dos resultados e net/http/httptest para simular requisições HTTP.

// O arquivo inclui três testes principais:
// 1. TestGetPaymentStatus_ValidRequest: Verifica se uma solicitação válida retorna o status correto da transação.
// 2. TestGetPaymentStatus_InvalidRequest_MissingTransactionID: Verifica se a ausência do transaction_id na solicitação resulta em um erro adequado.
// 3. TestGetPaymentStatus_InvalidRequest_UnsupportedGateway: Verifica se o uso de um gateway não suportado resulta em um erro adequado.

package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"desafiogolang-payment/handlers"
	"desafiogolang-payment/models"
	"desafiogolang-payment/services"

	"github.com/stretchr/testify/assert"
)

// Mock service response for GetPayPalPaymentStatus
func mockGetPayPalPaymentStatus(transactionID string) models.TransactionResponse {
	if transactionID == "valid-id" {
		return models.TransactionResponse{
			Message: "Transaction found",
			Status:  "completed",
		}
	}
	return models.TransactionResponse{
		Message: "Transaction ID not found",
		Status:  "unknown",
	}
}

func TestGetPaymentStatus_ValidRequest(t *testing.T) {
	// Injeção de dependência da função mock
	originalFunc := services.GetPayPalPaymentStatusFunc
	services.GetPayPalPaymentStatusFunc = mockGetPayPalPaymentStatus
	defer func() { services.GetPayPalPaymentStatusFunc = originalFunc }()

	// Cria uma solicitação válida
	req, err := http.NewRequest("GET", "/payment-status?transaction_id=valid-id&gateway=PayPal", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetPaymentStatus)

	// Chama o handler
	handler.ServeHTTP(rr, req)

	// Verifica o status da resposta
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verifica a resposta JSON
	var response models.TransactionResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Transaction found", response.Message)
	assert.Equal(t, "completed", response.Status)
}

func TestGetPaymentStatus_InvalidRequest_MissingTransactionID(t *testing.T) {
	// Cria uma solicitação sem o transaction_id
	req, err := http.NewRequest("GET", "/payment-status?gateway=PayPal", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetPaymentStatus)

	// Chama o handler
	handler.ServeHTTP(rr, req)

	// Verifica o status da resposta
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Transaction ID is required\n", rr.Body.String())
}

func TestGetPaymentStatus_InvalidRequest_UnsupportedGateway(t *testing.T) {
	// Cria uma solicitação com um gateway não suportado
	req, err := http.NewRequest("GET", "/payment-status?transaction_id=valid-id&gateway=Unknown", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetPaymentStatus)

	// Chama o handler
	handler.ServeHTTP(rr, req)

	// Verifica o status da resposta
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Unsupported gateway\n", rr.Body.String())
}
