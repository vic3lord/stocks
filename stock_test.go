package stocks

import (
	"testing"
)

func TestGetStock(t *testing.T) {
	stock := GetStock("aapl") //.List.Resources[0].Resource.Fields.Symbol
	if stock.GetSymbol(0) != "AAPL" {
		t.Error("The stock symbol should be equal to aapl but equals to", stock.GetSymbol(0))
	}
}
