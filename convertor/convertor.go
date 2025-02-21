package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func getExchangesRates(fromCurrency string, toCurrency string) error {
	parametrs := fromCurrency + "," + toCurrency
	fmt.Println("urparametrsl:", parametrs)
	url := fmt.Sprintf("https://api.frankfurter.dev/v1/latest?symbols=%s", parametrs)
	fmt.Println("url:", url)
	resp, err := http.Get(url)
	for {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[:n]))

		if n == 0 || err != nil {
			break
		}
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Print(parametrs)
	fmt.Print(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var rates string
	if err := json.Unmarshal(body, &rates); err != nil {
		return err
	}
	fmt.Print(rates)
	return err
}
func main() {
	// Исходная валюта
	fmt.Println("Введите исходную валюту (например USD)")
	var fromCurrency string
	fmt.Scanln(&fromCurrency)
	fromCurrency = strings.ToUpper(fromCurrency)
	fmt.Println("Введите сумму ", fromCurrency)

	// сумма валюты
	var countCurrencyStr string
	fmt.Scanln(&countCurrencyStr)
	countCurrencyStr = strings.ReplaceAll(countCurrencyStr, ",", ".")
	countCurrency, err := strconv.ParseFloat(countCurrencyStr, 64)
	if err != nil {
		fmt.Println("не число ", countCurrency)
		return
	}

	// Валюту в которую надо перевести
	fmt.Println("Введите Валюту в которую надо перевести (например RUB)")
	var toCurrency string
	fmt.Scanln(&toCurrency)
	toCurrency = strings.ToUpper(toCurrency)

	// функция получения курса валют
	getExchangesRates(fromCurrency, toCurrency)
}
