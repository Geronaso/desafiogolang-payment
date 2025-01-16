package services

import (
	"desafiogolang-payment/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetExchangeRate obtém a taxa de câmbio atual entre duas moedas.
func GetExchangeRate(fromCurrency, toCurrency string) (float64, error) {
	url := fmt.Sprintf("https://open.er-api.com/v6/latest/%s", fromCurrency)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	if result["result"] != "success" {
		return 0, fmt.Errorf("failed to get exchange rate")
	}

	rates := result["rates"].(map[string]interface{})
	toRate, exists := rates[toCurrency]
	if !exists {
		return 0, fmt.Errorf("currency not found")
	}

	return toRate.(float64), nil
}

// ConvertCurrency realiza a conversão de moeda usando a taxa de câmbio atual.
func ConvertCurrency(request models.CurrencyConversionRequest) (models.CurrencyConversionResponse, error) {
	rate, err := GetExchangeRate(request.FromCurrency, request.ToCurrency)
	if err != nil {
		return models.CurrencyConversionResponse{}, err
	}

	convertedAmount := request.Amount * rate
	return models.CurrencyConversionResponse{
		ConvertedAmount: convertedAmount,
		FromCurrency:    request.FromCurrency,
		ToCurrency:      request.ToCurrency,
		Rate:            rate,
	}, nil
}
