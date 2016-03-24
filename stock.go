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

// GetQuote - get full stock details into a struct
func GetQuote(symbol string) (Stock, error) {
	// set http client timeout
	client := http.Client{Timeout: timeout}

	url := fmt.Sprintf("http://finance.yahoo.com/webservice/v1/symbols/%s/quote?format=json&view=detail", symbol)
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

// LoopResults - return the stock name
func (stock Stock) LoopResults() {
	cnt := len(stock.List.Resources)
	for index, element := range stock.List.Resources {
		fmt.Printf("Index value [%d] of [%d] is [%s]\n", index+1, cnt, element)
	}
}

// GetName - return the stock name
func (stock Stock) GetName(index int) string {
	return stock.List.Resources[index].Resource.Fields.Name
}

// GetSymbol - return the stock symbol
func (stock Stock) GetSymbol(index int) string {
	return stock.List.Resources[index].Resource.Fields.Symbol
}

// GetPrice - return the stock price
func (stock Stock) GetPrice(index int) (float64, error) {
	price, err := strconv.ParseFloat(stock.List.Resources[index].Resource.Fields.Price, 64)
	if err != nil {
		return 1.0, fmt.Errorf("Stock price: %v", err)
	}
	return price, nil
}

// Get52WeekLow - return the 52 week low
func (stock Stock) Get52WeekLow(index int) (float64, error) {
	price, err := strconv.ParseFloat(stock.List.Resources[index].Resource.Fields.YearLow, 64)
	if err != nil {
		return 1.0, fmt.Errorf("52 week low: %v", err)
	}
	return price, nil
}

// Get52WeekHigh - return the 52 week high
func (stock Stock) Get52WeekHigh(index int) (float64, error) {
	price, err := strconv.ParseFloat(stock.List.Resources[index].Resource.Fields.YearHigh, 64)
	if err != nil {
		return 1.0, fmt.Errorf("52 week low: %v", err)
	}
	return price, nil
}

// GetDayLow - return the daily low
func (stock Stock) GetDayLow(index int) (float64, error) {
	price, err := strconv.ParseFloat(stock.List.Resources[index].Resource.Fields.DayLow, 64)
	if err != nil {
		return 1.0, fmt.Errorf("Daily low: %v", err)
	}
	return price, nil
}

// GetDayHigh - return the daily high
func (stock Stock) GetDayHigh(index int) (float64, error) {
	price, err := strconv.ParseFloat(stock.List.Resources[index].Resource.Fields.DayHigh, 64)
	if err != nil {
		return 1.0, fmt.Errorf("Daily high: %v", err)
	}
	return price, nil
}

// GetByIndex - return the 52 week high
func (stock Stock) GetByIndex(index int) string {
	price, err := stock.GetPrice(index)
	if err != nil {
		fmt.Printf("Error getting price: %v", err)
	}
	yearLow, err := stock.Get52WeekLow(index)
	if err != nil {
		fmt.Printf("Error getting 52weeklow: %v", err)
	}
	yearHigh, err := stock.Get52WeekHigh(index)
	if err != nil {
		fmt.Printf("Error getting 52weekHigh: %v", err)
	}
	dayLow, err := stock.GetDayLow(index)
	if err != nil {
		fmt.Printf("Error getting daily low: %v", err)
	}
	dayHigh, err := stock.GetDayHigh(index)
	if err != nil {
		fmt.Printf("Error getting daily high: %v", err)
	}
	return fmt.Sprintf("Name:\t%s\nSymbol:\t%s\nPrice:\t%f\nYearLow:\t%f\nYearHigh:\t%f\nDayLow:\t%f\nDayHigh:\t%f\n", stock.GetName(index), stock.GetSymbol(index), price, yearLow, yearHigh, dayLow, dayHigh)
}

// implement String()
func (stock Stock) String() string {
	return stock.GetByIndex(0)
}
