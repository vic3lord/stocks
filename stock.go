package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	timeout = time.Duration(time.Second * 10)
)

func GetStock(symbol string) StockType {
	client := http.Client{
		Timeout: timeout,
	}

	url := fmt.Sprintf("http://finance.yahoo.com/webservice/v1/symbols/%s/quote?format=json", symbol)
	res, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error connecting to API %q", err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading json body %q", err)
	}

	var stock StockType

	err = json.Unmarshal(content, &stock)
	if err != nil {
		fmt.Printf("Error reading json body %q", err)
	}

	return stock
}

func (stock StockType) GetSymbol(index int) string {
	return stock.List.Resources[index].Resource.Fields.Symbol
}
