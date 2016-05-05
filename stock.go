package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const timeout = time.Duration(time.Second * 10)

// GetQuote gets latest stock quote for a given symbol
func GetQuote(s string) (Stock, error) {
	var stock Stock
	client := http.Client{Timeout: timeout}

	api, err := url.Parse("https://query.yahooapis.com/v1/public/yql")
	if err != nil {
		return stock, fmt.Errorf("cannot parse yahooapis url: %v", err)
	}

	// Add query parameters
	params := url.Values{}
	params.Add("q", fmt.Sprintf("select * from yahoo.finance.quotes where symbol in (%q)", s))
	params.Add("format", "json")
	params.Add("env", "http://datatables.org/alltables.env")
	params.Add("callback", "")
	api.RawQuery = params.Encode()

	res, err := client.Get(api.String())
	if err != nil {
		return stock, fmt.Errorf("stocks cannot access yahoo finance API: %v", err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return stock, fmt.Errorf("stocks cannot read JSON body: %v", err)
	}

	err = json.Unmarshal(content, &stock)
	if err != nil {
		return stock, fmt.Errorf("stocks cannot parse JSON data: %v", err)
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
