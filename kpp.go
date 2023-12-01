package datrus

import (
	"fmt"

	"github.com/ofstudio/go-datrus/dict"
)

const lenKPP = 9

// KPP - КПП, код причины постановки на учет налогоплательщиков-организаций.
type KPP struct {
	Number        string         `json:"number"`         // 9-значный номер КПП.
	Region        dict.Region    `json:"region"`         // Регион КПП (первые 2 цифры номера КПП).
	AuthorityCode string         `json:"authority_code"` // Код налогового органа (первые 4 цифры номера КПП).
	Reason        dict.KPPReason `json:"reason"`         // Причина постановки на учет (5 и 6 цифры номера КПП).
	SerialNumber  string         `json:"serial_number"`  // Порядковый номер постановки на учёт в налоговом органе по соответствующему основанию (последние 3 цифры номера КПП).
}

// ParseKPP - анализирует строку с номером КПП и возвращает его детальную структуру.
// Входной формат: 9 цифр без разделителей и других символов.
// Возвращает ошибки ErrKPP и ErrLen или ErrChar, если строка не соответствует формату.
func ParseKPP(number string) (KPP, error) {
	if err := validateNumber(number, lenKPP); err != nil {
		return KPP{}, fmt.Errorf("%w: %w", ErrKPP, err)
	}
	return KPP{
		Number:        number,
		Region:        dict.RegionFindByCode(number[:2]),
		AuthorityCode: number[:4],
		Reason:        dict.KPPReasonFind(number[4:6]),
		SerialNumber:  number[6:9],
	}, nil
}

// MustParseKPP вызывает ParseKPP. В случае ошибки, завершается паникой.
func MustParseKPP(number string) KPP {
	kpp, err := ParseKPP(number)
	if err != nil {
		panic(err)
	}
	return kpp
}
