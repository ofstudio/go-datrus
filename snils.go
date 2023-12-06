package datrus

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
)

// SNILS - СНИЛС, страховой номер индивидуального лицевого счета, 11 цифр.
type SNILS string

const lenSNILS = 11

var reSNILS = regexp.MustCompile(`^(\d{3})[\s-]?(\d{3})[\s-]?(\d{3})[\s-]?(\d{2})$`)

// ParseSNILS - анализирует строку с номером СНИЛС и возвращает номер СНИЛС без дефисов и пробелов.
// Входной формат: 11 цифр, допускаются разделители пробелы и дефисы: "000-000-000 00".
// Возвращает ошибки:
//   - [ErrSNILS] и [ErrLen], если строка не соответствует формату
//   - [ErrSNILS] и [ErrCheckSum], если контрольная сумма не совпадает
func ParseSNILS(number string) (SNILS, error) {
	matches := reSNILS.FindStringSubmatch(number)
	if len(matches) != 5 {
		return "", fmt.Errorf("%w: %w", ErrSNILS, ErrLen)
	}

	number = matches[1] + matches[2] + matches[3] + matches[4]
	if len(number) != lenSNILS {
		return "", fmt.Errorf("%w: %w", ErrSNILS, ErrLen)
	}

	if number[9:11] != snilsCheckSum(number[:9]) {
		return "", fmt.Errorf("%w: %w", ErrSNILS, ErrCheckSum)
	}

	return SNILS(number), nil
}

// MustParseSNILS вызывает [ParseSNILS]. В случае ошибки, завершается паникой.
func MustParseSNILS(number string) SNILS {
	snils, err := ParseSNILS(number)
	if err != nil {
		panic(err)
	}
	return snils
}

const (
	snilsMin       = 1_001_999           // Минимальный номер СНИЛС: https://www.consultant.ru/document/cons_doc_LAW_142584/1d9155a863a5949b14b95ecbb536aa84856a2a2e/
	snilsMax       = 999_999_999         // Максимально возможный номер СНИЛС
	snilsRandLimit = snilsMax - snilsMin // Количество возможных вариантов номеров СНИЛС
)

// NewRandomSNILS создает новый СНИЛС со случайным номером в диапазоне
// от 001-001-999-xx до 999-999-999-xx включительно.
//
// Для генерации случайного номера используется crypto/rand.
// Если невозможно сгенерировать случайное значение,
// возвращает [ErrSNILS] и [ErrRand].
func NewRandomSNILS() (SNILS, error) {
	bigN, err := rand.Int(rand.Reader, big.NewInt(snilsRandLimit+1))
	if err != nil {
		return "", fmt.Errorf("%w: %w: %w", ErrSNILS, ErrRand, err)
	}
	n := fmt.Sprintf("%09d", bigN.Int64()+snilsMin)
	return SNILS(n + snilsCheckSum(n)), nil
}

// MustNewRandomSNILS вызывает [NewRandomSNILS]. В случае ошибки, завершается паникой.
func MustNewRandomSNILS() SNILS {
	snils, err := NewRandomSNILS()
	if err != nil {
		panic(err)
	}
	return snils
}

// Шаблоны для форматирования номера СНИЛС функцией [SNILS.Format].
const (
	SNILSDash      = "000-000-000-00"
	SNILSSpace     = "000 000 000 00"
	SNILSDashSpace = "000-000-000 00"
)

// Format форматирует номер СНИЛС в соответствии с шаблоном:
// [SNILSDash], [SNILSSpace] или [SNILSDashSpace].
func (s SNILS) Format(layout string) string {
	switch layout {
	case SNILSDashSpace:
		return fmt.Sprintf("%s-%s-%s %s", s[0:3], s[3:6], s[6:9], s[9:11])
	case SNILSDash:
		return fmt.Sprintf("%s-%s-%s-%s", s[0:3], s[3:6], s[6:9], s[9:11])
	case SNILSSpace:
		return fmt.Sprintf("%s %s %s %s", s[0:3], s[3:6], s[6:9], s[9:11])
	default:
		return string(s)
	}
}

func snilsCheckSum(num string) string {
	sum := 0
	for i, d := range num {
		sum += int(d-'0') * (9 - i)

	}
	sum = sum % 101
	if sum == 100 {
		sum = 0
	}
	return fmt.Sprintf("%02d", sum)
}
