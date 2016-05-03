package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	timeout  = time.Duration(time.Second * 10)
	urlStart = "https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20yahoo.finance.quotes%20where%20symbol%20in%20(%22"
	urlEnd   = "%22)%0A%09%09&format=json&env=http%3A%2F%2Fdatatables.org%2Falltables.env&callback="
)

// GetQuote gets latest stock quote for a given symbol
func GetQuote(symbol string) (Stock, error) {
	client := http.Client{Timeout: timeout}

	url := urlStart + symbol + urlEnd
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
	return s.Query.Results.Quote.Name
}

// GetSymbol gets stock symbol
func (s Stock) GetSymbol() string {
	return strings.ToUpper(s.Query.Results.Quote.Symbol)
}

// GetPrice gets the current stock price
func (s Stock) GetPrice() (float64, error) {
	p, err := strconv.ParseFloat(s.Query.Results.Quote.LastTradePriceOnly, 64)
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
