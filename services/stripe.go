// handlers.go
// Este arquivo é uma ideia inicial para implementação de outro gateway de pagamento.
// Mas a ideia seria a mesma do arquivo paypal.go. Deverá ser validado a API do Stripe e implementado aqui a comunicação com a mesma.

package services

import "desafiogolang-payment/models"

func ProcessStripePayment(request models.PaymentRequest) models.PaymentResponse {
	return models.PaymentResponse{
		Transaction_ID: "success",
		Message:        "Payment processed successfully",
	}
}
