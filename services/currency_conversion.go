// currency_conversion.go
// Este arquivo contém funções para realizar a conversão de moeda e obter taxas de câmbio atualizadas.
// Utiliza uma estrutura de cache para armazenar temporariamente as taxas de câmbio e evitar consultas excessivas à API externa.

// O arquivo inclui duas funções principais e uma variável de função mockável:
// 1. GetExchangeRate: Obtém a taxa de câmbio atual entre duas moedas, utilizando cache para armazenar as taxas mais recentes.
// 2. convertCurrency: Realiza a conversão de moeda usando a taxa de câmbio atual.
// 3. ConvertCurrencyFunc: Variável de função mockável que permite substituir a implementação da função convertCurrency durante os testes.

package services

import (
	"desafiogolang-payment/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type exchangeRateCache struct {
	rates      map[string]map[string]float64
	timestamps map[string]time.Time
	mu         sync.Mutex
}

var cache = exchangeRateCache{
	rates:      make(map[string]map[string]float64),
	timestamps: make(map[string]time.Time),
}

// GetExchangeRate obtém a taxa de câmbio atual entre duas moedas.
func GetExchangeRate(fromCurrency, toCurrency string) (float64, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	// Verifica se a taxa está no cache e se é da última hora
	if rates, ok := cache.rates[fromCurrency]; ok {
		if rate, exists := rates[toCurrency]; exists {
			if time.Since(cache.timestamps[fromCurrency]) < time.Hour {
				return rate, nil
			}
		}
	}

	// Se não estiver no cache ou estiver desatualizada, consulta a API externa
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

	// Atualiza o cache
	if cache.rates[fromCurrency] == nil {
		cache.rates[fromCurrency] = make(map[string]float64)
	}
	cache.rates[fromCurrency][toCurrency] = toRate.(float64)
	cache.timestamps[fromCurrency] = time.Now()

	return toRate.(float64), nil
}

// Mockable function variable
var ConvertCurrencyFunc = convertCurrency

// convertCurrency realiza a conversão de moeda usando a taxa de câmbio atual.
func convertCurrency(request models.CurrencyConversionRequest) (models.CurrencyConversionResponse, error) {
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

func ConvertCurrency(request models.CurrencyConversionRequest) (models.CurrencyConversionResponse, error) {
	return ConvertCurrencyFunc(request)
}
