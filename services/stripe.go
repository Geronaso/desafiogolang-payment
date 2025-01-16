package services

import "desafiogolang-payment/models"

func ProcessStripePayment(request models.PaymentRequest) models.PaymentResponse {
	// Implemente a lógica de processamento de pagamento do PayPal aqui
	return models.PaymentResponse{
		Transaction_ID: "success",
		Message:        "Payment processed successfully",
	}
}
