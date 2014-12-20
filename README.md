# Stocks

Stocks is a go library to get stock by symbol in from yahoo finance  

## Example

```
$ go get github.com/vic3lord/stocks
```

```go
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
	fmt.Printf("Google stock price is: %f", stock.GetPrice(0))
	stock.PrettyPrint()
}
```

## Disclaimer

Just had a free time to hack something and so I did this was not tested carefully  
**DO NOT USE IN PROD**

