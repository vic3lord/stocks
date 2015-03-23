package stocks

import "fmt"

func ExampleGetSymbol() {
	stock, _ := GetQuote("aapl")
	fmt.Println(stock.GetSymbol())
	// Output: AAPL
}

func ExampleGetName() {
	stock, _ := GetQuote("aapl")
	fmt.Println(stock.GetName())
	// Output: Apple Inc.
}
