# Stocks

Stocks is a go library to get stock by symbol in from yahoo finance  

[![GoDoc](https://godoc.org/github.com/vic3lord/stocks?status.svg)](https://godoc.org/github.com/vic3lord/stocks)

## Example

```
$ go get github.com/vic3lord/stocks
```

**Example 1**

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
	fmt.Print(stock)
}
```

**Example 2**

```go
package main

import (
	"fmt"
	"github.com/vic3lord/stocks"
	"os"
	"os/signal"
)

var input string

func GrabStock() {
	fmt.Print("> ")
	fmt.Scanf("%s\n", &input)
}

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)

	go func() {
		<-exit
		fmt.Println("Exiting...")
		os.Exit(0)
	}()

	fmt.Println("Press CTRL-C to exit!\nChoose stock symbol to get quote:")

	for {
		GrabStock()
		stock, err := stocks.GetQuote(input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(stock)
	}
}
```

## ISC License

```
Copyright (c) 2015, Or Elimelech <0r3limelech@gmail.com>

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
```
