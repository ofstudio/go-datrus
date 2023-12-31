package dict

var currencyIdxCode = make(map[string]*Currency)   // Индекс по цифровому коду.
var currencyIdxSymbol = make(map[string]*Currency) // Индекс по буквенному коду.

// Шортркаты валют
var (
	CurrencyRUR = Currency{ /* init() */ } // Российский рубль
	CurrencyEUR = Currency{ /* init() */ } // Евро
	CurrencyUSD = Currency{ /* init() */ } // Доллар США
)

func init() {
	// Построение индексов
	for i, c := range currencyDict {
		currencyIdxCode[c.Code] = &currencyDict[i]
		currencyIdxSymbol[c.Symbol] = &currencyDict[i]
	}

	// Инициализация шорткатов
	CurrencyRUR = CurrencyFindByCode("810")
	CurrencyEUR = CurrencyFindByCode("978")
	CurrencyUSD = CurrencyFindByCode("840")
}

// currencyDict - справочник кодов валют.
// Цифровые и буквенные коды, наименования валют в соответствии с
// "ОК (МК (ИСО 4217) 003-97) 014-2000. Общероссийский классификатор валют"
//
// Источник: https://www.cbr.ru/development/mcirabis/kv/
//
// Актуальность справочника: 2023-08-25.
var currencyDict = []Currency{
	{"978", "EUR", "ЕВРО"},
	{"826", "GBP", "ФУНТ СТЕРЛИНГОВ"},
	{"036", "AUD", "АВСТРАЛИЙСКИЙ ДОЛЛАР"},
	{"703", "SKK", "СЛОВАЦКАЯ КРОНА"},
	{"716", "ZWD", "ДОЛЛАР ЗИМБАБВЕ"},
	{"566", "NGN", "НАЙРА"},
	{"532", "ANG", "НИДЕРЛАНДСКИЙ АНТИЛЬСКИЙ ГУЛЬДЕН"},
	{"795", "TMM", "МАНАТ"},
	{"762", "TJR", "ТАДЖИКСКИЙ РУБЛ"},
	{"233", "EEK", "КРОНА"},
	{"008", "ALL", "ЛЕК"},
	{"946", "RON", "НОВЫЙ РУМЫНСКИЙ ЛЕЙ"},
	{"096", "BND", "БРУНЕЙСКИЙ ДОЛЛАР"},
	{"116", "KHR", "РИЕЛЬ"},
	{"736", "SDD", "СУДАНСКИЙ ДИНАР"},
	{"484", "MXN", "МЕКСИКАНСКОЕ ПЕСО"},
	{"152", "CLP", "ЧИЛИЙСКОЕ ПЕСО"},
	{"554", "NZD", "НОВОЗЕЛАНДСКИЙ ДОЛЛАР"},
	{"752", "SEK", "ШВЕДСКАЯ КРОНА"},
	{"203", "CZK", "ЧЕШСКАЯ КРОНА"},
	{"953", "XPF", "ФРАНК КФП"},
	{"840", "USD", "ДОЛЛАР США"},
	{"935", "ZWR", "ДОЛЛАР ЗИМБАБВЕ"},
	{"634", "QAR", "КАТАРСКИЙ РИАЛ"},
	{"360", "IDR", "РУПИЯ"},
	{"050", "BDT", "ТАКА"},
	{"478", "MRO", "УГИЯ"},
	{"440", "LTL", "ЛИТОВСКИЙ ЛИТ"},
	{"328", "GYD", "ГАЙАНСКИЙ ДОЛЛАР"},
	{"810", "RUR", "РОССИЙСКИЙ РУБЛЬ"},
	{"174", "KMF", "КОМОРСКИЙ ФРАНК"},
	{"941", "RSD", "СЕРБСКИЙ ДИНАР"},
	{"970", "COU", "ЕДИНИЦА РЕАЛЬНОЙ СТОИМОСТИ"},
	{"408", "KPW", "СЕВЕРО-КОРЕЙСКАЯ ВОНА"},
	{"516", "NAD", "ДОЛЛАР НАМИБИИ"},
	{"480", "MUR", "МАВРИКИЙСКАЯ РУПИЯ"},
	{"020", "ADP", "АНДОРРСКАЯ ПЕСЕТА"},
	{"616", "PLZ", "ЗЛОТЫЙ"},
	{"496", "MNT", "ТУГРИК"},
	{"214", "DOP", "ДОМИНИКАНСКОЕ ПЕСО"},
	{"524", "NPR", "НЕПАЛЬСКАЯ РУПИЯ"},
	{"368", "IQD", "ИРАКСКИЙ ДИНАР"},
	{"724", "ESP", "ИСПАНСКАЯ ПЕСЕТА"},
	{"450", "MGF", "МАЛАГАСИЙСКИЙ ФРАНК"},
	{"678", "STD", "ДОБРА"},
	{"858", "UYP", "УРУГВАЙСКОЕ ПЕСО"},
	{"072", "BWP", "ПУЛА"},
	{"860", "UZS", "УЗБЕКСКИЙ СУМ"},
	{"124", "CAD", "КАНАДСКИЙ ДОЛЛАР"},
	{"952", "XOF", "ФРАНК КФА ВСЕАО"},
	{"882", "WST", "ТАЛА"},
	{"780", "TTD", "ДОЛЛАР ТРИНИДАДА И ТОБАГО"},
	{"710", "ZAR", "РЭНД"},
	{"418", "LAK", "КИП"},
	{"937", "VEF", "БОЛИВАР"},
	{"422", "LBP", "ЛИВАНСКИЙ ФУНТ"},
	{"654", "SHP", "ФУНТ ОСТРОВА СВЯТОЙ ЕЛЕНЫ"},
	{"188", "CRC", "КОСТАРИКАНСКИЙ КОЛОН"},
	{"238", "FKP", "ФУНТ ФОЛКЛЕНДСКИХ ОСТРОВОВ"},
	{"404", "KES", "КЕНИЙСКИЙ ШИЛЛИНГ"},
	{"886", "YER", "ЙЕМЕНСКИЙ РИАЛ"},
	{"728", "SSP", "ЮЖНОСУДАНСКИЙ ФУНТ"},
	{"246", "FIM", "МАРКА"},
	{"690", "SCR", "СЕЙШЕЛЬСКАЯ РУПИЯ"},
	{"012", "DZD", "АЛЖИРСКИЙ ДИНАР"},
	{"348", "HUF", "ФОРИНТ"},
	{"558", "NIO", "ЗОЛОТАЯ КОРДОБА"},
	{"068", "BOB", "БОЛИВИАНО"},
	{"960", "XDR", "СДР (СПЕЦИАЛЬНЫЕ ПРАВА ЗАИМСТВОВАНИЯ)"},
	{"414", "KWD", "КУВЕЙТСКИЙ ДИНАР"},
	{"760", "SYP", "СИРИЙСКИЙ ФУНТ"},
	{"944", "AZN", "АЗЕРБАЙДЖАНСКИЙ МАНАТ"},
	{"498", "MDL", "МОЛДАВСКИЙ ЛЕЙ"},
	{"230", "ETB", "ЭФИОПСКИЙ БЫР"},
	{"938", "SDG", "СУДАНСКИЙ ФУНТ"},
	{"974", "BYR", "БЕЛОРУССКИЙ РУБЛЬ"},
	{"454", "MWK", "КВАЧА"},
	{"292", "GIP", "ГИБРАЛТАРСКИЙ ФУНТ"},
	{"818", "EGP", "ЕГИПЕТСКИЙ ФУНТ"},
	{"376", "ILS", "ШЕКЕЛЬ"},
	{"044", "BSD", "БАГАМСКИЙ ДОЛЛАР"},
	{"590", "PAB", "БАЛЬБОА"},
	{"090", "SBD", "ДОЛЛАР СОЛОМОНОВЫХ ОСТРОВОВ"},
	{"032", "ARS", "АРГЕНТИНСКОЕ ПЕСО"},
	{"528", "NLG", "НИДЕРЛАНДСКИЙ ГУЛЬДЕН"},
	{"784", "AED", "ДИРХАМ (ОАЭ)"},
	{"807", "MKD", "ДЕНАР"},
	{"694", "SLL", "ЛЕОНЕ"},
	{"586", "PKR", "ПАКИСТАНСКАЯ РУПИЯ"},
	{"426", "LSL", "ЛОТИ"},
	{"894", "ZMK", "КВАЧА (ЗАМБИЙСКАЯ)"},
	{"004", "AFA", "АФГАНИ"},
	{"929", "MRU", "УГИЯ"},
	{"976", "CDF", "КОНГОЛЕЗСКИЙ ФРАНК"},
	{"643", "RUB", "РОССИЙСКИЙ РУБЛЬ"},
	{"458", "MYR", "МАЛАЙЗИЙСКИЙ РИНГГИТ"},
	{"470", "MTL", "МАЛЬТИЙСКАЯ ЛИРА"},
	{"156", "CNY", "ЮАНЬ РЕНМИНБИ"},
	{"985", "PLN", "ЗЛОТЫЙ"},
	{"950", "XAF", "ФРАНК КФА ВЕАС"},
	{"398", "KZT", "ТЕНГЕ"},
	{"340", "HNL", "ЛЕМПИРА"},
	{"136", "KYD", "ДОЛЛАР КАЙМАНОВЫХ ОСТРОВОВ"},
	{"951", "XCD", "ВОСТОЧНО-КАРИБСКИЙ ДОЛЛАР"},
	{"604", "PEN", "НОВЫЙ СОЛЬ"},
	{"980", "UAH", "ГРИВНА"},
	{"967", "ZMW", "ЗАМБИЙСКАЯ КВАЧА"},
	{"108", "BIF", "БУРУНДИЙСКИЙ ФРАНК"},
	{"800", "UGX", "УГАНДИЙСКИЙ ШИЛЛИНГ"},
	{"364", "IRR", "ИРАНСКИЙ РИАЛ"},
	{"400", "JOD", "ИОРДАНСКИЙ ДИНАР"},
	{"642", "ROL", "СТАРЫЙ РУМЫНСКИЙ ЛЕЙ"},
	{"945", "AYM", "АЗЕРБАЙДЖАНСКИЙ МАНАТ"},
	{"232", "ERN", "НАКФА"},
	{"646", "RWF", "ФРАНК РУАНДЫ"},
	{"288", "GHC", "СЕДИ"},
	{"943", "MZN", "МОЗАМБИКСКИЙ МЕТИКАЛ"},
	{"936", "GHS", "СЕДИ"},
	{"901", "TWD", "НОВЫЙ ТАЙВАНЬСКИЙ ДОЛЛАР"},
	{"270", "GMD", "ДАЛАСИ"},
	{"048", "BHD", "БАХРЕЙНСКИЙ ДИНАР"},
	{"064", "BTN", "НГУЛТРУМ"},
	{"934", "TMT", "НОВЫЙ МАНАТ"},
	{"446", "MOP", "ПАТАКА"},
	{"144", "LKR", "ШРИ-ЛАНКИЙСКАЯ РУПИЯ"},
	{"410", "KRW", "ВОНА"},
	{"834", "TZS", "ТАНЗАНИЙСКИЙ ШИЛЛИНГ"},
	{"191", "HRK", "ХОРВАТСКАЯ КУНА"},
	{"975", "BGN", "БОЛГАРСКИЙ ЛЕВ"},
	{"862", "VEB", "БОЛИВАР"},
	{"430", "LRD", "ЛИБЕРИЙСКИЙ ДОЛЛАР"},
	{"344", "HKD", "ГОНКОНГСКИЙ ДОЛЛАР"},
	{"702", "SGD", "СИНГАПУРСКИЙ ДОЛЛАР"},
	{"973", "AOA", "КВАНЗА"},
	{"792", "TRL", "СТАРАЯ ТУРЕЦКАЯ ЛИРА"},
	{"608", "PHP", "ФИЛИППИНСКОЕ ПЕСО"},
	{"428", "LVL", "ЛАТВИЙСКИЙ ЛАТ"},
	{"890", "YUN", "ЮГОСЛАВСКИЙ ДИНАР"},
	{"276", "DEM", "НЕМЕЦКАЯ МАРКА"},
	{"626", "TPE", "ТИМОРСКОЕ ЭСКУДО"},
	{"196", "CYP", "КИПРСКИЙ ФУНТ"},
	{"740", "SRG", "СУРИНАМСКИЙ ГУЛЬДЕН"},
	{"392", "JPY", "ИЕНА"},
	{"084", "BZD", "БЕЛИЗСКИЙ ДОЛЛАР"},
	{"052", "BBD", "БАРБАДОССКИЙ ДОЛЛАР"},
	{"442", "LUF", "ЛЮКСЕМБУРГСКИЙ ФРАНК"},
	{"508", "MZM", "МЕТИКАЛ"},
	{"192", "CUP", "КУБИНСКОЕ ПЕСО"},
	{"748", "SZL", "ЛИЛАНГЕНИ"},
	{"100", "BGL", "ЛЕВ"},
	{"218", "ECS", "СУКРЕ"},
	{"891", "YUM", "НОВЫЙ ДИНАР"},
	{"380", "ITL", "ИТАЛЬЯНСКАЯ ЛИРА"},
	{"417", "KGS", "СОМ (КИРГИЗСКИЙ)"},
	{"977", "BAM", "КОНВЕРТИРУЕМАЯ МАРКА"},
	{"986", "BRL", "БРАЗИЛЬСКИЙ РЕАЛ"},
	{"051", "AMD", "АРМЯНСКИЙ ДРАМ"},
	{"704", "VND", "ДОНГ"},
	{"931", "CUC", "КОНВЕРТИРУЕМОЕ ПЕСО"},
	{"056", "BEF", "БЕЛЬГИЙСКИЙ ФРАНК"},
	{"949", "TRY", "ТУРЕЦКАЯ ЛИРА"},
	{"504", "MAD", "МАРОККАНСКИЙ ДИРХАМ"},
	{"262", "DJF", "ФРАНК ДЖИБУТИ"},
	{"112", "BYB", "БЕЛОРУССКИЙ РУБЛЬ"},
	{"933", "BYN", "БЕЛОРУССКИЙ РУБЛЬ"},
	{"208", "DKK", "ДАТСКАЯ КРОНА"},
	{"512", "OMR", "ОМАНСКИЙ РИАЛ"},
	{"170", "COP", "КОЛУМБИЙСКОЕ ПЕСО"},
	{"434", "LYD", "ЛИВИЙСКИЙ ДИНАР"},
	{"132", "CVE", "ЭСКУДО КАБО-ВЕРДЕ"},
	{"705", "SIT", "ТОЛАР"},
	{"972", "TJS", "СОМОНИ"},
	{"324", "GNF", "ГВИНЕЙСКИЙ ФРАНК"},
	{"764", "THB", "БАТ"},
	{"969", "MGA", "АРИАРИ"},
	{"578", "NOK", "НОРВЕЖСКАЯ КРОНА"},
	{"104", "MMK", "КЬЯТ"},
	{"756", "CHF", "ШВЕЙЦАРСКИЙ ФРАНК"},
	{"971", "AFN", "АФГАНИ"},
	{"940", "UYI", "УРУГВАЙСКОЕ ПЕСО В ИНДЕКСИРОВАННЫХ ЕДИНИЦАХ"},
	{"776", "TOP", "ПААНГА"},
	{"548", "VUV", "ВАТУ"},
	{"932", "ZWL", "ДОЛЛАР ЗИМБАБВЕ"},
	{"352", "ISK", "ИСЛАНДСКАЯ КРОНА"},
	{"031", "AZM", "АЗЕРБАЙДЖАНСКИЙ МАНАТ"},
	{"280", "DEM", "НЕМЕЦКАЯ МАРКА"},
	{"600", "PYG", "ГУАРАНИ"},
	{"250", "FRF", "ФРАНЦУЗСКИЙ ФРАНК"},
	{"620", "PTE", "ПОРТУГАЛЬСКОЕ ЭСКУДО"},
	{"682", "SAR", "САУДОВСКИЙ РИЯЛ"},
	{"388", "JMD", "ЯМАЙСКИЙ ДОЛЛАР"},
	{"180", "ZRN", "НОВЫЙ ЗАИР"},
	{"533", "AWG", "АРУБАНСКИЙ ГУЛЬДЕН"},
	{"060", "BMD", "БЕРМУДСКИЙ ДОЛЛАР"},
	{"242", "FJD", "ДОЛЛАР ФИДЖИ"},
	{"268", "GEK", "ГРУЗИНСКИЙ КУПОН"},
	{"024", "AON", "НОВАЯ КВАНЗА"},
	{"222", "SVC", "САЛЬВАДОРСКИЙ КОЛОН"},
	{"981", "GEL", "ЛАРИ"},
	{"040", "ATS", "ШИЛЛИНГ"},
	{"968", "SRD", "СУРИНАМСКИЙ ДОЛЛАР"},
	{"300", "GRD", "ДРАХМА"},
	{"598", "PGK", "КИНА"},
	{"372", "IEP", "ИРЛАНДСКИЙ ФУНТ"},
	{"930", "STN", "ДОБРА"},
	{"706", "SOS", "СОМАЛИЙСКИЙ ШИЛЛИНГ"},
	{"332", "HTG", "ГУРД"},
	{"320", "GTQ", "КЕТСАЛЬ"},
	{"462", "MVR", "РУФИЯ"},
	{"624", "GWP", "ПЕСО ГВИНЕИ-БИСАУ"},
	{"356", "INR", "ИНДИЙСКАЯ РУПИЯ"},
	{"788", "TND", "ТУНИССКИЙ ДИНАР"},
	{"942", "ZWN", "НОВЫЙ ДОЛЛАР ЗИМБАБВЕ"},
	{"954", "XEU", "ЭКЮ (ЕДИНИЦА ЕВРОПЕЙСКОЙ ВАЛЮТЫ)"},
	{"928", "VES", "БОЛИВАР СОБЕРАНО"},
	{"925", "SLE", "ЛЕОНЕ"},
	{"926", "VED", "БОЛИВАР СОБЕРАНО"},
}
