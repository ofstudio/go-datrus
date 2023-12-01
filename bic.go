package datrus

import "fmt"

// http://cbr.ru/PSystem/payment_system/

const lenBIC = 9

// BICParticipantType - тип участника платежной системы.
type BICParticipantType string

const (
	BICDirectParticipant   BICParticipantType = "Участник платежной системы с прямым участием"
	BICInDirectParticipant BICParticipantType = "Участник платежной системы с косвенным участием"
	BICNonParticipant      BICParticipantType = "Клиент Банка России, не являющийся участником платежной системы"
	BICUnknownParticipant  BICParticipantType = "Неизвестный вид участия"
)

// BIC - БИК, банковский идентификационный код.
type BIC struct {
	Number          string             `json:"number"`           // 9-значный номер БИК.
	ParticipantType BICParticipantType `json:"participant_type"` // Информация об участии и виде участия в платежной системе.
}

// ParseBIC - анализирует строку с номером БИК и возвращает его детальную структуру.
// Входной формат: 9 цифр без разделителей и других символов.
// Возвращает ошибки ErrBIC и ErrLen или ErrChar, если строка не соответствует формату.
func ParseBIC(number string) (BIC, error) {
	if err := validateNumber(number, lenBIC); err != nil {
		return BIC{}, fmt.Errorf("%w: %w", ErrBIC, err)
	}

	var participantType BICParticipantType
	switch number[0] {
	case '0':
		participantType = BICDirectParticipant
	case '1':
		participantType = BICInDirectParticipant
	case '2':
		participantType = BICNonParticipant
	default:
		participantType = BICUnknownParticipant
	}

	return BIC{Number: number, ParticipantType: participantType}, nil
}

// MustParseBIC вызывает ParseBIC. В случае ошибки, завершается паникой.
func MustParseBIC(number string) BIC {
	account, err := ParseBIC(number)
	if err != nil {
		panic(err)
	}
	return account
}

// CheckBankAccount - проверяет соответствие ("ключевку") номера расчетного счета и БИК.
// В случае несоответствия возвращает ошибки ErrBIC и ErrBICCheck.
func (b BIC) CheckBankAccount(account BankAccount) error {
	if err := validateBICAccount(b.Number[6:9] + account.Number); err != nil {
		return fmt.Errorf("%w: %w", ErrBIC, err)
	}
	return nil
}

// CheckCorrAccount - проверяет соответствие ("ключевку") номера корреспондентского счета и БИК.
// В случае несоответствия возвращает ошибки ErrBIC и ErrBICCheck.
func (b BIC) CheckCorrAccount(account BankAccount) error {
	if err := validateBICAccount("0" + b.Number[4:6] + account.Number); err != nil {
		return fmt.Errorf("%w: %w", ErrBIC, err)
	}
	return nil
}

var bicAccountWeights = []int32{7, 1, 3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7, 1}

func validateBICAccount(input string) error {
	runes := []rune(input)
	if len(runes) != len(bicAccountWeights) {
		return ErrBICCheck
	}

	var sum int32
	for pos, char := range input {
		if char < '0' || char > '9' {
			return ErrBICCheck
		}
		d := char - '0'
		sum += (bicAccountWeights[pos] * d) % 10
	}

	if sum%10 != 0 {
		return ErrBICCheck
	}

	return nil
}
