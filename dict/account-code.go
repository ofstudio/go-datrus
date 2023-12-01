package dict

//go:generate go run ../internal/dict-account-code-gen/main.go -in account-code-dict.txt -out account-code-dict.go

func init() {
	// Код 30101 исключен из текущего перечня, тем не менее применяется.
	// Указание Банка России от 24 марта 1998 года № 191-У
	// https://normativ.kontur.ru/document?moduleId=1&documentId=28266
	accountDict["30101"] = "Корреспондентские счета кредитных организаций в Банке России"

	// todo Уточнить этот и другие ранее исключенные коды счетов
	accountDict["30103"] = "Корреспондентские счета расчетных небанковских кредитных организаций"
}

// AccountCode - код счета первого или второго порядка.
type AccountCode struct {
	Code  string `json:"code"`  // Код счета (3 цифры для счета первого порядка, 5 цифр для счетов второго порядка)
	Title string `json:"title"` // Наименование кода счета
}

const AccountCodeUnknown = "Неизвестный код счета" // AccountCode.Title неизвестного кода счета.

// IsUnknown возвращает true для неизвестного кода счета.
func (c AccountCode) IsUnknown() bool {
	return c.Title == AccountCodeUnknown
}

// AccountCodeFind - возвращает наименование балансового счета
// по его коду первого или второго порядка согласно Положению №385-П ЦБ РФ
// "О правилах ведения бухгалтерского учета в кредитных организациях, расположенных на территории Российской Федерации".
// Если код не найден в справочнике, возвращает код с AccountCode.Code = AccountCodeUnknown.
//
// Источники:
//   - https://cbr.ru/Queries/UniDbQuery/File/85920?fileId=-1&scope=1374-1375
//   - https://www.profbanking.com/chart-of-accounts-in-banks
//
// Актуальность справочника: 2023-05-13.
func AccountCodeFind(code string) AccountCode {
	if title, ok := accountDict[code]; ok {
		return AccountCode{Code: code, Title: title}
	}
	return AccountCode{Code: code, Title: AccountCodeUnknown}
}
