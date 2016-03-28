package main

import (
	"fmt"

	"github.com/vic3lord/stocks"
)

func main() {
	stock, err := stocks.GetQuote("goog")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(stock)
}
