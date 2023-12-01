package datrus

import "errors"

// Ошибки первого уровня
var (
	ErrBankAccount = errors.New("ошибка банковского счета")
	ErrBIC         = errors.New("ошибка БИК")
	ErrINN         = errors.New("ошибка ИНН")
	ErrKPP         = errors.New("ошибка КПП")
	ErrSNILS       = errors.New("ошибка SNILS")
)

// Ошибки второго уровня
var (
	ErrLen      = errors.New("неверная длина")
	ErrChar     = errors.New("недопустимый символ")
	ErrCheckSum = errors.New("ошибка контрольной суммы")
	ErrBICCheck = errors.New("счет не соответствует БИК")
	ErrRand     = errors.New("ошибка генерации случайного значения")
)
