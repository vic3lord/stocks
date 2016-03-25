package main

import (
	"fmt"

	"github.com/sgravitz/stocks"
)

func main() {
	stock, err := stocks.GetQuote("fb,goog,ge")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(stock)
	fmt.Print(stock.GetByIndex(1))
	stock.LoopResults()

}
