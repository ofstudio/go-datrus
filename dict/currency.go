package dict

import "strings"

// Currency - код и наименование валюты.
type Currency struct {
	Code   string `json:"code"`   // Цифровой код валюты. Пример: "840".
	Symbol string `json:"symbol"` // Буквенный код валюты. Пример: "USD".
	Title  string `json:"title"`  // Наименование валюты. Пример: "ДОЛЛАР США".
}

const CurrencyUnknown = "Неизвестная валюта" // Currency.Title  неизвестной валюты.

// IsUnknown - возвращает true для неизвестной валюты
func (c Currency) IsUnknown() bool {
	return c.Title == CurrencyUnknown
}

// CurrencyFindByCode ищет валюту по ее цифровому коду в соответствии с
// "ОК (МК (ИСО 4217) 003-97) 014-2000. Общероссийский классификатор валют".
// Если валюта не найдена, вернет валюту с Currency.Title = CurrencyUnknown.
//
// Источники
//   - https://www.cbr.ru/development/mcirabis/kv/
//
// Актуальность справочника: 2023-08-25.
func CurrencyFindByCode(code string) Currency {
	if c, ok := currencyIdxCode[code]; ok {
		return *c
	}
	return Currency{Code: code, Title: CurrencyUnknown}
}

// CurrencyFindBySymbol ищет валюту по ее буквенному коду в соответствии с
// "ОК (МК (ИСО 4217) 003-97) 014-2000. Общероссийский классификатор валют".
// Если валюта не найдена, вернет валюту с Currency.Title = CurrencyUnknown.
//
// Источники
//   - https://www.cbr.ru/development/mcirabis/kv/
//
// Актуальность справочника: 2023-08-25.
func CurrencyFindBySymbol(symbol string) Currency {
	if c, ok := currencyIdxSymbol[strings.ToUpper(symbol)]; ok {
		return *c
	}
	return Currency{Symbol: symbol, Title: CurrencyUnknown}
}
