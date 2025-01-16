package models

// CurrencyConversionRequest representa uma solicitação de conversão de moeda.
type CurrencyConversionRequest struct {
	Amount       float64 `json:"amount" validate:"required,gt=0"`
	FromCurrency string  `json:"from_currency" validate:"required,len=3"`
	ToCurrency   string  `json:"to_currency" validate:"required,len=3"`
}

// CurrencyConversionResponse representa a resposta de uma conversão de moeda.
type CurrencyConversionResponse struct {
	ConvertedAmount float64 `json:"converted_amount"`
	FromCurrency    string  `json:"from_currency"`
	ToCurrency      string  `json:"to_currency"`
	Rate            float64 `json:"rate"`
}
