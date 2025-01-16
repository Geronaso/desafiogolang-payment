//paypal.go
// Este módulo SIMULA de maneira BEM SIMPLIFICADA uma transação para a API do paypal.
// É gerado um ID de transação e um status de pagamento de maneira aleatório e guardado em memória.
// A outra função checa e retorna o status baseado nesse ID.
// Idealmente deveria ser implementado utilizando a API de sandbox do Paypal e utilizar os endpoints comentados, levando em consideração TODAS as variaveis descritas e acesso a banco de dados.
//https://developer.paypal.com/docs/api/payments/v1/#payment_get
//https://developer.paypal.com/docs/api/payments/v1/#payment_create

package services

import (
	"desafiogolang-payment/models"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Gera as var
var (
	transactions     = make(map[string]models.Transaction)
	transactionsLock sync.Mutex
	rng              = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// generateTransactionID gera um ID de transação único
func generateTransactionID() string {
	return fmt.Sprintf("PAY-%d", rng.Intn(1000000000))
}

// ProcessPayPalPayment simula o processamento de um pagamento no PayPal
func ProcessPayPalPayment(request models.PaymentRequest) models.PaymentResponse {
	transactionID := generateTransactionID()

	// Simulando diferentes resultados com base em valores aleatórios
	statuses := []string{"completed", "pending", "failed"}
	status := statuses[rng.Intn(len(statuses))]

	transactionsLock.Lock()
	transactions[transactionID] = models.Transaction{
		Status:         status,
		Transaction_ID: transactionID,
	}
	transactionsLock.Unlock()

	return models.PaymentResponse{
		Message:        "Payment processed with success",
		Transaction_ID: transactionID,
	}
}

// GetPayPalPaymentStatus simula a verificação do status de uma transação no PayPal
func GetPayPalPaymentStatus(transactionID string) models.TransactionResponse {
	transactionsLock.Lock()
	defer transactionsLock.Unlock()

	if transaction, exists := transactions[transactionID]; exists {
		return models.TransactionResponse{
			Message: fmt.Sprintf("Transaction ID: %s found", transactionID),
			Status:  transaction.Status,
		}
	}
	return models.TransactionResponse{
		Message: "Transaction ID not found",
		Status:  "unknown",
	}
}
