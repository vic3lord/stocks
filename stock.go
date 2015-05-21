package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	timeout = time.Duration(time.Second * 10)
)

// get full stock details into a struct
func GetQuote(symbol string) (Stock, error) {
	// set http client timeout
	client := http.Client{Timeout: timeout}

	url := fmt.Sprintf("http://finance.yahoo.com/webservice/v1/symbols/%s/quote?format=json", symbol)
	res, err := client.Get(url)
	if err != nil {
		return Stock{}, fmt.Errorf("Stocks cannot access yahoo finance API: %v", err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Stock{}, fmt.Errorf("Stocks cannot read json body: %v", err)
	}

	var stock Stock

	err = json.Unmarshal(content, &stock)
	if err != nil {
		return Stock{}, fmt.Errorf("Stocks cannot parse json data: %v", err)
	}

	return stock, nil
}

// return the stock name
func (stock Stock) GetName() string {
	return stock.List.Resources[0].Resource.Fields.Name
}

// return the stock symbol
func (stock Stock) GetSymbol() string {
	return stock.List.Resources[0].Resource.Fields.Symbol
}

// return the stock price
func (stock Stock) GetPrice() (float64, error) {
	price, err := strconv.ParseFloat(stock.List.Resources[0].Resource.Fields.Price, 64)
	if err != nil {
		return 1.0, fmt.Errorf("Stock price: %v", err)
	}

	return price, nil
}

// just print all details nicely
func (stock Stock) PrettyPrint() {
	name := stock.GetName()
	sym := stock.GetSymbol()
	price, err := stock.GetPrice()
	if err != nil {
		fmt.Printf("Error getting price: %v", err)
	}

	fmt.Printf("Name:\t%s\nSymbol:\t%s\nPrice:\t%f\n", name, sym, price)
}
