package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/vic3lord/stocks"
)

var input string

// GrabStock - read stock symbol from input
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
