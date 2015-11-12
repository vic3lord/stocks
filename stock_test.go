package stocks

import "fmt"

func ExampleStock_GetSymbol() {
	stock, err := GetQuote("aapl")
	if err != nil {
		fmt.Printf("Error getting quote: %v", err)
	}
	fmt.Println(stock.GetSymbol())
	// Output: AAPL
}

func ExampleStock_GetName() {
	stock, err := GetQuote("aapl")
	if err != nil {
		fmt.Printf("Error getting quote: %v", err)
	}
	fmt.Println(stock.GetName())
	// Output: Apple Inc.
}
