package datrus

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/ofstudio/go-datrus/dict"
)

func ExampleParseBankAccount() {
	acc, _ := ParseBankAccount("40817810745256078055")
	fmt.Println(acc.L2.Title)
	fmt.Println(acc.Currency.Title)

	// Output: Физические лица
	// РОССИЙСКИЙ РУБЛЬ

}

func ExampleParseBankAccount_error() {
	_, err := ParseBankAccount("4081781074525607abcd")
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, ErrBankAccount), errors.Is(err, ErrChar))
	}

	// Output: ошибка банковского счета: недопустимый символ
	// true true
}

func TestParseBankAccount(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		want    BankAccount
		wantErr error
	}{
		{
			name:  "успешно",
			input: "30101810500000000641",
			want: BankAccount{
				Number:        "30101810500000000641",
				L1:            dict.AccountCode{Code: "301", Title: "Корреспондентские счета"},
				L2:            dict.AccountCode{Code: "30101", Title: "Корреспондентские счета кредитных организаций в Банке России"},
				Currency:      dict.Currency{Code: "810", Symbol: "RUR", Title: "РОССИЙСКИЙ РУБЛЬ"},
				DivisionCode:  "0000",
				AccountNumber: "0000641",
			},
		},
		{
			name:    "недопустимый символ",
			input:   "1234567890123invalid",
			wantErr: ErrChar,
		},
		{
			name:    "неверная длина",
			input:   "123",
			wantErr: ErrLen,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBankAccount(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ParseBankAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseBankAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
