package stocks

import "testing"

var testData = []struct {
	in     string
	name   string
	symbol string
}{
	{"aapl", "Apple Inc.", "AAPL"},
	{"goog", "Alphabet Inc.", "GOOG"},
}

func TestGetSymbol(t *testing.T) {
	for _, tt := range testData {
		s, err := GetQuote(tt.in)
		if err != nil {
			t.Errorf("Error getting quote for %s: %v", tt.in, err)
		}

		got := s.GetSymbol()
		if got != tt.symbol {
			t.Errorf("GetSymbol(%q) = %s; want %s", tt.in, got, tt.symbol)
		}
	}
}

func TestGetName(t *testing.T) {
	for _, tt := range testData {
		s, err := GetQuote(tt.in)
		if err != nil {
			t.Errorf("Error getting quote for %s: %v", tt.in, err)
		}

		got := s.GetName()
		if got != tt.name {
			t.Errorf("GetSymbol(%q) = %s; want %s", tt.in, got, tt.name)
		}
	}
}
