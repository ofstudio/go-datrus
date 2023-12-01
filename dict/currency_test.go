package dict

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCurrencyShortcuts(t *testing.T) {
	tests := []struct {
		name     string
		currency Currency
		want     string
	}{
		{"CurrencyRUR", CurrencyRUR, "RUR"},
		{"CurrencyEUR", CurrencyEUR, "EUR"},
		{"CurrencyUSD", CurrencyUSD, "USD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.currency.Symbol; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s.Symbol = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestCurrencyFindByCode(t *testing.T) {
	tests := []struct {
		name string
		code string
		want Currency
	}{
		{
			name: "успешно",
			code: "840",
			want: Currency{Code: "840", Symbol: "USD", Title: "ДОЛЛАР США"},
		},
		{
			name: "не найдено",
			code: "000",
			want: Currency{Code: "000", Title: CurrencyUnknown},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CurrencyFindByCode(tt.code)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrencyFindByCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrencyFindBySymbol(t *testing.T) {

	tests := []struct {
		name   string
		symbol string
		want   Currency
	}{
		{
			name:   "успешно",
			symbol: "USD",
			want:   Currency{Code: "840", Symbol: "USD", Title: "ДОЛЛАР США"},
		},
		{
			name:   "не найдено",
			symbol: "$$$",
			want:   Currency{Symbol: "$$$", Title: CurrencyUnknown},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CurrencyFindBySymbol(tt.symbol); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrencyFindBySymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrency_IsUnknown(t *testing.T) {

	tests := []struct {
		name     string
		currency Currency
		want     bool
	}{
		{
			name:     "известная валюта",
			currency: CurrencyRUR,
			want:     false,
		},
		{
			name:     "неизвестный код",
			currency: CurrencyFindByCode("0"),
			want:     true,
		},
		{
			name:     "неизвестный символ",
			currency: CurrencyFindBySymbol("$$$"),
			want:     true,
		},
		{
			name: "неизвестное наименование",
			currency: Currency{
				Code:   "000",
				Symbol: "AAA",
				Title:  CurrencyUnknown,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.currency.IsUnknown(); got != tt.want {
				t.Errorf("IsUnknown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCurrencyFindByCode() {
	fmt.Println(CurrencyFindByCode("810"))
	fmt.Println(CurrencyFindByCode("840"))
	fmt.Println(CurrencyFindByCode("999").Title)

	// Output: {810 RUR РОССИЙСКИЙ РУБЛЬ}
	// {840 USD ДОЛЛАР США}
	// Неизвестная валюта

}

func ExampleCurrencyFindBySymbol() {
	fmt.Println(CurrencyFindBySymbol("RUR"))
	fmt.Println(CurrencyFindBySymbol("eur"))
	fmt.Println(CurrencyFindBySymbol("XXX").Title)

	// Output: {810 RUR РОССИЙСКИЙ РУБЛЬ}
	// {978 EUR ЕВРО}
	// Неизвестная валюта
}
