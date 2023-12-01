package dict

import "strings"

const (
	lenOKATO = 11       // Полная длина кода ОКАТО
	lenOKTMO = lenOKATO // Полная длина кода ОКТМО
)

// Region - регион (субъект) РФ.
type Region struct {
	// Код региона согласно справочника ФНС "Справочник субъектов Российской Федерации".
	// Пример: "02" (Республика Башкортостан)
	Code string

	// Наименование региона.
	// Пример: "Алтайский край".
	Title string

	// Код ОКАТО региона (2 или 5 цифр).
	// Примеры: "05" (Приморский край),"71140" (Ямало-Ненецкий автономный округ)
	OKATO string

	// Код ОКТМО региона (2 или 5 цифр).
	// Примеры: "05" (Приморский край),"71900" (Ямало-Ненецкий автономный округ)
	OKTMO string
}

const RegionUnknown = "Неизвестный регион" // Region.Title неизвестного региона.

// IsUnknown возвращает true для неизвестного региона
func (r Region) IsUnknown() bool {
	return r.Title == RegionUnknown
}

// RegionFindByCode - возвращает регион по коду согласно
// справочника ФНС "Справочник субъектов Российской Федерации".
// Если регион не найден, вернет регион с Region.Title = RegionUnknown.
//
// Источники:
//
//   - https://www.nalog.gov.ru/rn77/program/5961292/
//
// Актуальность справочника: 2022-10-05.
func RegionFindByCode(code string) Region {
	if r, ok := regionIdxCode[code]; ok {
		return *r
	}
	return Region{Code: code, Title: RegionUnknown}
}

// RegionFindByOKATO - возвращает регион по его коду ОКАТО.
// Если регион не найден, вернет регион с Region.Title = RegionUnknown.
//
// Источники:
//
//   - https://rosstat.gov.ru/opendata/7708234640-7708234640-okato
//
// Актуальность справочника: 2023-11-01
func RegionFindByOKATO(code string) Region {
	if r, ok := regionIdxOKATO[normalizeTerrCode(code, lenOKATO)]; ok {
		return *r
	}
	return Region{OKATO: code, Title: RegionUnknown}
}

// RegionFindByOKTMO - возвращает регион по его коду ОКТМО.
// Если регион не найден, вернет регион с Region.Title = RegionUnknown.
//
// Источники:
//
//   - https://rosstat.gov.ru/opendata/7708234640-oktmo
//
// Актуальность справочника: 2023-11-01
func RegionFindByOKTMO(code string) Region {
	if r, ok := regionIdxOKTMO[normalizeTerrCode(code, lenOKTMO)]; ok {
		return *r
	}
	return Region{OKTMO: code, Title: RegionUnknown}
}

// normalizeTerrCode - заполняет нулями недостающие разряды кодов ОКАТО и ОКТМО.
func normalizeTerrCode(code string, l int) string {
	if n := l - len([]rune(code)); n > 0 {
		code += strings.Repeat("0", n)
	}
	return code
}
