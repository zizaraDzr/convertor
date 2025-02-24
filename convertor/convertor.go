package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type ExchangeRates struct {
	Amount float64            `json:"amount"`
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}

func getExchangesRates(fromCurrency string, toCurrency string) (*ExchangeRates, error) {
	url := fmt.Sprintf("https://api.frankfurter.dev/v1/latest?base=%s&symbols=%s", fromCurrency, toCurrency)
	fmt.Println("url:", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка: сервер вернул статус %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения тела ответа: %w", err)
	}
	var result ExchangeRates
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("ошибка разбора JSON: %w", err)
	}
	return &result, nil
}
func main() {
	// Исходная валюта
	fmt.Print("Соотношение евро к другим валютам\n")
	fmt.Println("Введите исходную валюту (например USD)")
	var fromCurrency string
	fmt.Scanln(&fromCurrency)
	fromCurrency = strings.ToUpper(fromCurrency)

	// Валюту в которую надо перевести
	fmt.Println("Введите Валюту в которую надо перевести (например GBP)")
	var toCurrency string
	fmt.Scanln(&toCurrency)
	toCurrency = strings.ToUpper(toCurrency)

	// сумма валюты
	fmt.Println("Введите сумму ")
	var countCurrency string
	fmt.Scanln(&countCurrency)
	countCurrency = strings.ReplaceAll(countCurrency, ",", ".")
	amount, err := strconv.ParseFloat(countCurrency, 64)
	if err != nil {
		fmt.Println("не число ", amount)
		return
	}
	// функция получения курса валют
	rates, err := getExchangesRates(fromCurrency, toCurrency)
	if err != nil {
		fmt.Println("ошибка в получении курса валют ", err)
		return
	}
	res, found := rates.Rates[toCurrency]
	if !found {
		fmt.Println("нет такой валюты ", toCurrency)
		return
	}
	res = amount * res

	fmt.Printf("%g %s = %.2f %s", amount, fromCurrency, res, toCurrency)
}
