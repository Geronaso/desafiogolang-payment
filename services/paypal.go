// paypal.go
// Este módulo simula de maneira simplificada uma transação para a API do PayPal.
// Ele gera um ID de transação e um status de pagamento de maneira aleatória e guarda essas informações em memória.
// Outra função checa e retorna o status baseado nesse ID.

// Idealmente, a implementação deveria utilizar a API de sandbox do PayPal e os endpoints comentados abaixo,
// levando em consideração todas as variáveis descritas, tratamento de erros da API externa e por fim criado uma camada de repositório para acesso ao banco de dados.
// https://developer.paypal.com/docs/api/payments/v1/#payment_get
// https://developer.paypal.com/docs/api/payments/v1/#payment_create

package services

import (
	"desafiogolang-payment/models"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	transactions     = make(map[string]models.Transaction)
	transactionsLock sync.Mutex
	rng              = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Mockable function variable
var GetPayPalPaymentStatusFunc = getPayPalPaymentStatus

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

// getPayPalPaymentStatus simula a verificação do status de uma transação no PayPal
func getPayPalPaymentStatus(transactionID string) models.TransactionResponse {
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

// GetPayPalPaymentStatus é feito para efeitos de testes.
// Como os dados são armazenados em memória, foi necessário atribuir a uma variável a função,
// permitindo modificar e mockar o seu comportamento durante os testes.
// Idealmente, utilizando um banco de dados para armazenamento dos dados, isso não seria necessário.
func GetPayPalPaymentStatus(transactionID string) models.TransactionResponse {
	return GetPayPalPaymentStatusFunc(transactionID)
}
