package dict

// KPPReason - причина постановки на учет налогоплательщиков-организаций.
type KPPReason struct {
	Code  string `json:"code"`  // Двузначный код причины (5 и 6 цифры номера КПП).
	Title string `json:"title"` // Наименование причины.
}

const KPPReasonUnknown = "Неизвестная причина" // KPPReason.Title неизвестной причины

// KPPReasonFind - возвращает причину постановки на учет в соответствии
// со Справочником "Причины постановки на учет налогоплательщиков-организаций" (СППУНО).
// Если код не найден в справочнике, возвращает причину с KPPReason.Title = [KPPReasonUnknown].
//
// Источник:
//   - https://normativ.kontur.ru/document?moduleId=1&documentId=244201
//     с дополнениями согласно Письма ФНС от 2 июня 2008 г. N ЧД-6-6/396
//     «О применении кодов справочника "СППУНО"» (коды 43, 44, 45).
//
// Актуальность справочника: 2008-06-02 (?)
func KPPReasonFind(code string) KPPReason {
	if r, ok := kppReasonDict[code]; ok {
		return r
	}
	return KPPReason{Code: code, Title: KPPReasonUnknown}
}

// IsUnknown возвращает true для неизвестной причины.
func (r KPPReason) IsUnknown() bool {
	return r.Title == KPPReasonUnknown
}
