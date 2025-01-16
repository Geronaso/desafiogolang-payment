package handlers

import (
	"desafiogolang-payment/models"
	"desafiogolang-payment/services" // Importando o pacote services
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Inicializa uma instância do validador
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ProcessPayment lida com solicitações de pagamento, decodificando a solicitação JSON,
// validando os dados e encaminhando para o gateway de pagamento apropriado.
func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var paymentRequest models.PaymentRequest

	// Decodifica o corpo da solicitação JSON em uma estrutura PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&paymentRequest); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Valida a estrutura paymentRequest
	if err := validate.Struct(paymentRequest); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Processa o pagamento com base no gateway especificado
	switch paymentRequest.Gateway {
	case "PayPal":
		// Chama a função processPayPalPayment do pacote services para processar o pagamento
		response := services.ProcessPayPalPayment(paymentRequest)
		// Codifica a resposta em JSON e envia de volta ao cliente
		json.NewEncoder(w).Encode(response)
	case "Stripe":
		// Chama a função processPayPalPayment do pacote services para processar o pagamento
		response := services.ProcessStripePayment(paymentRequest)
		// Codifica a resposta em JSON e envia de volta ao cliente
		json.NewEncoder(w).Encode(response)
	default:
		// Retorna um erro se o gateway não for suportado
		http.Error(w, "Unsupported gateway", http.StatusBadRequest)
	}
}

// GetPaymentStatus lida com solicitações para verificar o status de uma transação
func GetPaymentStatus(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transaction_id")
	payment_gateway := r.URL.Query().Get("gateway")
	if transactionID == "" {
		http.Error(w, "Transaction ID is required", http.StatusBadRequest)
		return
	}
	switch payment_gateway {
	case "PayPal":
		// Obtém o status da transação do serviço
		response := services.GetPayPalPaymentStatus(transactionID)
		// Codifica a resposta em JSON e envia de volta ao cliente
		json.NewEncoder(w).Encode(response)
	default:
		// Retorna um erro se o gateway não for suportado
		http.Error(w, "Unsupported gateway", http.StatusBadRequest)
	}
}
