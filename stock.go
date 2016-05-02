package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const timeout = time.Duration(time.Second * 10)

// GetQuote gets latest stock quote for a given symbol
func GetQuote(symbol string) (Stock, error) {
	client := http.Client{Timeout: timeout}

	url := fmt.Sprintf("http://finance.yahoo.com/webservice/v1/symbols/%s/quote?format=json", symbol)
	res, err := client.Get(url)
	if err != nil {
		return Stock{}, fmt.Errorf("stocks cannot access yahoo finance API: %v", err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Stock{}, fmt.Errorf("stocks cannot read JSON body: %v", err)
	}

	var stock Stock
	err = json.Unmarshal(content, &stock)
	if err != nil {
		return Stock{}, fmt.Errorf("stocks cannot parse JSON data: %v", err)
	}
	return stock, nil
}

// GetName gets stock name
func (s Stock) GetName() string {
	return s.List.Resources[0].Resource.Fields.Name
}

// GetSymbol gets stock symbol
func (s Stock) GetSymbol() string {
	return s.List.Resources[0].Resource.Fields.Symbol
}

// GetPrice gets the current stock price
func (s Stock) GetPrice() (float64, error) {
	p, err := strconv.ParseFloat(s.List.Resources[0].Resource.Fields.Price, 64)
	if err != nil {
		return 1.0, fmt.Errorf("stock price: %v", err)
	}

	return p, nil
}

// String implements String()
func (s Stock) String() string {
	p, err := s.GetPrice()
	if err != nil {
		fmt.Printf("Error getting price: %v", err)
	}

	return fmt.Sprintf("Name:\t%s\nSymbol:\t%s\nPrice:\t%f\n", s.GetName(), s.GetSymbol(), p)
}
