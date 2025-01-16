// payment.go
// Este arquivo define os modelos de dados necessários para o sistema de pagamentos
// e inclui as respectivas restrições de validação para garantir a integridade dos dados.

package models

// PaymentRequest representa uma solicitação de pagamento.
// Inclui detalhes do gateway, valor, moeda (somente USD é aceito), método de pagamento e informações do cartão.
type PaymentRequest struct {
	Gateway       string      `json:"gateway" validate:"required"`
	Amount        float64     `json:"amount" validate:"required,gt=0"`
	Currency      string      `json:"currency" validate:"required,oneof=USD"`
	PaymentMethod string      `json:"payment_method" validate:"required"`
	CardDetails   CardDetails `json:"card_details" validate:"required"`
}

// CardDetails representa os detalhes do cartão de crédito.
type CardDetails struct {
	Number string `json:"number" validate:"required,len=16"`
	Expiry string `json:"expiry" validate:"required,len=5"`
	CVV    string `json:"cvv" validate:"required,len=3"`
}

// PaymentResponse representa a resposta de uma transação de pagamento.
type PaymentResponse struct {
	Message        string `json:"message"`
	Transaction_ID string `json:"transaction_id"`
}

type TransactionResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Transaction representa a estrutura de dados de uma transação interna
type Transaction struct {
	Status         string `json:"status"`
	Transaction_ID string `json:"message"`
}
