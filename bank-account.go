package datrus

import (
	"fmt"

	"github.com/ofstudio/go-datrus/dict"
)

// BankAccount - детальная структура банковского счета.
type BankAccount struct {
	Number        string           `json:"number"`         // 20-значный номер счета
	L1            dict.AccountCode `json:"l1"`             // Код счета первого порядка (первые 3 цифры номера счета)
	L2            dict.AccountCode `json:"l2"`             // Код счета второго порядка (первые 5 цифр номера счета)
	Currency      dict.Currency    `json:"currency"`       // Валюта счета
	DivisionCode  string           `json:"division_code"`  // Код филиала
	AccountNumber string           `json:"account_number"` // Номер лицевого счета
}

// ParseBankAccount - анализирует строку с номером банковского счета
// и возвращает его детальную структуру.
// Входной формат: 20 цифр без разделителей и других символов.
// Возвращает ошибку ErrBankAccount и ErrLen или ErrChar, если строка не соответствует формату.
func ParseBankAccount(number string) (BankAccount, error) {
	if err := validateNumber(number, 20); err != nil {
		return BankAccount{}, fmt.Errorf("%w: %w", ErrBankAccount, err)
	}

	return BankAccount{
		Number:        number,
		L1:            dict.AccountCodeFind(number[0:3]),
		L2:            dict.AccountCodeFind(number[0:5]),
		Currency:      dict.CurrencyFindByCode(number[5:8]),
		DivisionCode:  number[9:13],
		AccountNumber: number[13:20],
	}, nil
}

// MustParseBankAccount вызывает ParseBankAccount.
// В случае ошибки, завершается паникой.
func MustParseBankAccount(input string) BankAccount {
	account, err := ParseBankAccount(input)
	if err != nil {
		panic(err)
	}
	return account
}
