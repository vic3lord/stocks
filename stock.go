package stocks

import (
	"encoding/json"
	"errors"
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
func GetQuote(symbol string) (StockType, error) {
	client := http.Client{
		Timeout: timeout,
	}

	url := fmt.Sprintf("http://finance.yahoo.com/webservice/v1/symbols/%s/quote?format=json", symbol)
	res, err := client.Get(url)
	if err != nil {
		return StockType{}, errors.New("Stocks cannot access yahoo finance API...")
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return StockType{}, errors.New("Stocks cannot read json body...")
	}

	var stock StockType

	err = json.Unmarshal(content, &stock)
	if err != nil {
		return StockType{}, errors.New("Stocks cannot parse json data...")
	}

	return stock, nil
}

// return the stock name
func (stock StockType) GetName() string {
	return stock.List.Resources[0].Resource.Fields.Name
}

// return the stock symbol
func (stock StockType) GetSymbol() string {
	return stock.List.Resources[0].Resource.Fields.Symbol
}

// return the stock price
func (stock StockType) GetPrice() float64 {
	price, err := strconv.ParseFloat(stock.List.Resources[0].Resource.Fields.Price, 64)
	if err != nil {
		fmt.Println(err)
	}

	return price
}

// just print all details nicely
func (stock StockType) PrettyPrint() {
	fmt.Println("-------------------------------")
	fmt.Println("Name:\t", stock.GetName())
	fmt.Println("Symbol:\t", stock.GetSymbol())
	fmt.Println("Price:\t", stock.GetPrice())
	fmt.Println("-------------------------------")
}
