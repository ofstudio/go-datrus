package datrus

import (
	"fmt"

	"github.com/ofstudio/go-datrus/dict"
)

// INNEntityType - принадлежность ИНН к ФЛ или ЮЛ.
type INNEntityType string

const (
	INNPersonal  INNEntityType = "ФЛ" // ИНН физического лица.
	INNJuridical INNEntityType = "ЮЛ" // ИНН юридического лица.
)

const (
	lenINNPersonal  = 12 // Длина ИНН для ФЛ.
	lenINNJuridical = 10 // Длина ИНН для ЮЛ.
)

// INN - ИНН, индивидуальный номер налогоплательщика.
type INN struct {
	Number        string        `json:"number"`         // Номер ИНН: 12 цифр для ФЛ и 10 цифр для ЮЛ.
	EntityType    INNEntityType `json:"entity_type"`    // Субъект ИНН: INNPersonal ("ФЛ") или INNJuridical ("ЮЛ").
	AuthorityCode string        `json:"authority_code"` // Код налогового органа (первые 4 цифры ИНН).
	Region        dict.Region   `json:"region"`         // Регион налогоплательщика.
}

// ParseINN - анализирует строку с номером ИНН и возвращает его детальную структуру.
// Входной формат: 10 (для ЮЛ) или 12 цифр (для ФЛ) без разделителей и других символов.
// Возвращает ошибки:
//   - ErrINN и ErrLen или ErrChar, если строка не соответствует формату
//   - ErrINN и ErrCheckSum, если контрольная сумма не совпадает
func ParseINN(number string) (INN, error) {
	var entityType INNEntityType
	var err error

	switch len([]rune(number)) {
	case lenINNJuridical:
		entityType = INNJuridical
		err = validateINNJuridical(number)
	case lenINNPersonal:
		entityType = INNPersonal
		err = validateINNPersonal(number)
	default:
		err = ErrLen
	}

	if err != nil {
		return INN{}, fmt.Errorf("%w: %w", ErrINN, err)
	}

	return INN{
		Number:        number,
		EntityType:    entityType,
		AuthorityCode: number[0:4],
		Region:        dict.RegionFindByCode(number[0:2]),
	}, nil
}

// MustParseINN вызывает ParseINN. В случае ошибки, завершается паникой.
func MustParseINN(number string) INN {
	inn, err := ParseINN(number)
	if err != nil {
		panic(err)
	}
	return inn
}

var innWeights = []int32{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

func innSum(runes []rune, weights []int32) int32 {
	var s int32
	for pos, char := range runes {
		s += (char - '0') * weights[pos]
	}
	return s % 11 % 10
}

func validateINNPersonal(number string) error {

	if err := validateNumber(number, lenINNPersonal); err != nil {
		return err
	}
	runes := []rune(number)

	// Check sum1
	sum := innSum(runes[:lenINNPersonal-2], innWeights[1:])
	if runes[lenINNPersonal-2]-'0' != sum {
		return ErrCheckSum
	}

	// Check sum2
	sum = innSum(runes[:lenINNPersonal-1], innWeights)
	if runes[lenINNPersonal-1]-'0' != sum {
		return ErrCheckSum
	}

	return nil
}

func validateINNJuridical(number string) error {

	if err := validateNumber(number, lenINNJuridical); err != nil {
		return err
	}
	runes := []rune(number)

	// Check sum
	sum := innSum(runes[:lenINNJuridical-1], innWeights[2:])
	if runes[lenINNJuridical-1]-'0' != sum {
		return ErrCheckSum
	}

	return nil
}
