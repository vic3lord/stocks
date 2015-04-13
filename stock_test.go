package stocks

import "fmt"

func ExampleGetSymbol() {
	stock, err := GetQuote("aapl")
	if err != nil {
		fmt.Errorf("Error getting quote: %v", err)
	}
	fmt.Println(stock.GetSymbol())
	// Output: AAPL
}

func ExampleGetName() {
	stock, err := GetQuote("aapl")
	if err != nil {
		fmt.Errorf("Error getting quote: %v", err)
	}
	fmt.Println(stock.GetName())
	// Output: Apple Inc.
}
