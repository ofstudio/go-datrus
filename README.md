# go-datrus

[![Go Reference](https://pkg.go.dev/badge/github.com/ofstudio/go-datrus.svg)](https://pkg.go.dev/github.com/ofstudio/go-datrus)

`datrus` — библиотека для работы с различными идентификаторами и реквизитами, принятыми в РФ.

## Типы данных

### Банковский счет

[BankAccount](https://pkg.go.dev/github.com/ofstudio/go-datrus#BankAccount) — детальная структура банковского счета.
- Валидация
- Разбор структуры банковского счета
- Обогащение информацией о счетах первого и второго порядка по справочнику 
- Обогащение информацией о валюте по справочнику

### БИК

[BIC](https://pkg.go.dev/github.com/ofstudio/go-datrus#BIC) — БИК, банковский идентификационный код.
- Валидация
- Разбор структуры БИК
- Проверка соответствия ("ключевка") номеров расчетных и корреспондентских счеты с БИК

### ИНН

[INN](https://pkg.go.dev/github.com/ofstudio/go-datrus#INN) — ИНН, индивидуальный номер налогоплательщика.

- Валидация ИНН физических и юридических лиц
- Проверка контрольного числа
- Разбор структуры ИНН
- Обогащение информацией о регионе по справочнику

### КПП

[KPP](https://pkg.go.dev/github.com/ofstudio/go-datrus#KPP) — КПП, код причины постановки на учет 
налогоплательщиков-организаций.

- Валидация
- Разбор структуры КПП
- Обогащение информацией о регионе по справочнику
- Обогащение информацией о причине постановки на учет по справочнику


### СНИЛС 

[SNILS](https://pkg.go.dev/github.com/ofstudio/go-datrus#SNILS) — СНИЛС, страховой номер индивидуального лицевого счета в системе обязательного пенсионного страхования.

- Валидация
- Проверка контрольного числа
- Форматирование номера СНИЛС по шаблону
- Генерация случайного номера СНИЛС

## Справочники

`datrus/dict` — различные справочники с возможностью поиска. 
Источник данных и дата актуальности справочников указаны описаниях функций поиска 
по соответствующему справочнику.

 - [AccountCode](https://pkg.go.dev/github.com/ofstudio/go-datrus/dict#AccountCode) —
   справочник кодов счетов первого и второго порядка ("план счетов") согласно Положению №385-П ЦБ РФ
   "О правилах ведения бухгалтерского учета в кредитных организациях, расположенных
   на территории Российской Федерации"
 - [Currency](https://pkg.go.dev/github.com/ofstudio/go-datrus/dict#Currency) — справочник кодов валют 
   "ОК (МК (ИСО 4217) 003-97) 014-2000. Общероссийский классификатор валют"
 - [KPPReason](https://pkg.go.dev/github.com/ofstudio/go-datrus/dict#KPPReason) — справочник 
   "Причины постановки на учет налогоплательщиков-организаций" (СППУНО)
 - [Region](https://pkg.go.dev/github.com/ofstudio/go-datrus/dict#Region) — справочник субъектов 
   Российской Федерации с кодами ОКАТО и ОКТМО.

## Требования

- Go 1.21+

## Установка

```
go get -u github.com/ofstudio/go-datrus
```

## Примеры

### Соответствие БИК и номера банковского счета

```go
package main

import (
  "fmt"

  "github.com/ofstudio/go-datrus"
)

func main() {
  bic := datrus.MustParseBIC("044525700")

  // Счет  соответствует БИК
  account := datrus.MustParseBankAccount("40817810901004428532")
  if err := bic.CheckBankAccount(account); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Счет соответствует БИК")
  }

  // Счет не соответствует БИК
  account = datrus.MustParseBankAccount("40817810901004428533")
  if err := bic.CheckBankAccount(account); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Счет соответствует БИК")
  }

}
```
Вывод:
```
Счет соответствует БИК
ошибка БИК: счет не соответствует БИК
```

### Поиск по справочникам

```go
package main

import (
	"fmt"

	"github.com/ofstudio/go-datrus/dict"
)

func main() {
	region := dict.RegionFindByCode("89")
	fmt.Println(region)

	currency := dict.CurrencyFindBySymbol("USD")
	fmt.Println(currency)
}
```

Вывод:
```
{89 Ямало-Ненецкий автономный округ 71140000000 71900000000}
{840 USD ДОЛЛАР США}
```

## Лицензия
Распространяется по лицензии MIT. Более подробная информация в файле LICENSE.
