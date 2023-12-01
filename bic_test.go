package datrus

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func ExampleParseBIC() {
	bic, _ := ParseBIC("044525225")
	fmt.Println(bic)

	// Output: {044525225 Участник платежной системы с прямым участием}

}
func ExampleParseBIC_error() {
	_, err := ParseBIC("1234")
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, ErrBIC), errors.Is(err, ErrLen))
	}

	// Output: ошибка БИК: неверная длина
	// true true

}

func ExampleBIC_CheckBankAccount() {
	bic := MustParseBIC("044525700")
	acc1 := MustParseBankAccount("40817810901004428532")
	acc2 := MustParseBankAccount("40817810901004428533")

	if err := bic.CheckBankAccount(acc1); err != nil {
		fmt.Printf("Счет %s НЕ соответствует БИК %s: %v\n", acc1.Number, bic.Number, err)
	} else {
		fmt.Printf("Счет %s соответствует БИК %s\n", acc1.Number, bic.Number)
	}

	if err := bic.CheckBankAccount(acc2); err != nil {
		fmt.Printf("Счет %s НЕ соответствует БИК %s: %v\n", acc2.Number, bic.Number, err)
	} else {
		fmt.Printf("Счет %s соответствует БИК %s\n", acc2.Number, bic.Number)
	}

	// Output: Счет 40817810901004428532 соответствует БИК 044525700
	// Счет 40817810901004428533 НЕ соответствует БИК 044525700: ошибка БИК: счет не соответствует БИК

}

func TestBIC_CheckBankAccount(t *testing.T) {
	tests := []struct {
		name    string
		account BankAccount
		wantErr error
	}{
		{name: "успешно", account: MustParseBankAccount("40817810901004428532")},
		{name: "не соответствует счет", account: MustParseBankAccount("40817810901004428533"), wantErr: ErrBICCheck},
		{name: "неверная длина", account: BankAccount{Number: "1234"}, wantErr: ErrBICCheck},
		{name: "недопустимый символ", account: BankAccount{Number: "4081781090100442853x"}, wantErr: ErrBICCheck},
	}
	bic := MustParseBIC("044525700")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := bic.CheckBankAccount(tt.account); !errors.Is(err, tt.wantErr) {
				t.Errorf("CheckBankAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBIC_CheckCorrAccount(t *testing.T) {
	tests := []struct {
		name    string
		account BankAccount
		wantErr error
	}{
		{name: "успешно", account: MustParseBankAccount("30101810200000000700")},
		{name: "не соответствует счет", account: MustParseBankAccount("30101810200000000701"), wantErr: ErrBICCheck},
		{name: "неверная длина", account: BankAccount{Number: "1234"}, wantErr: ErrBICCheck},
		{name: "недопустимый символ", account: BankAccount{Number: "3010181020000000070x"}, wantErr: ErrBICCheck},
	}
	bic := MustParseBIC("044525700")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := bic.CheckCorrAccount(tt.account); !errors.Is(err, tt.wantErr) {
				t.Errorf("CheckCorrAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseBIC(t *testing.T) {

	tests := []struct {
		name    string
		number  string
		want    BIC
		wantErr error
	}{
		{name: "неверная длина", number: "12345678", wantErr: ErrLen},
		{name: "недопустимый символ", number: "1234x5678", wantErr: ErrChar},
		{name: "успешно", number: "044525225", want: BIC{Number: "044525225", ParticipantType: BICDirectParticipant}},
		{name: "неизвестный тип", number: "944525225", want: BIC{Number: "944525225", ParticipantType: BICUnknownParticipant}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBIC(tt.number)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ParseBIC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseBIC() got = %v, want %v", got, tt.want)
			}
		})
	}
}
