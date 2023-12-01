package datrus

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/ofstudio/go-datrus/dict"
)

func ExampleParseINN() {
	inn, _ := ParseINN("1653001805")
	fmt.Println(inn.EntityType)
	fmt.Println(inn.Region.Title)

	// Output: ЮЛ
	// Республика Татарстан (Татарстан)

}

func ExampleParseINN_error() {
	_, err := ParseINN("1653001806")
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, ErrINN), errors.Is(err, ErrCheckSum))
	}

	// Output: ошибка ИНН: ошибка контрольной суммы
	// true true

}

func TestParseINN(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name    string
		number  string
		want    INN
		wantErr error
	}{
		{
			name:   "успешно ФЛ",
			number: "772170926460",
			want: INN{
				Number:        "772170926460",
				EntityType:    INNPersonal,
				AuthorityCode: "7721",
				Region:        dict.Region{Code: "77", Title: "город федерального значения Москва", OKATO: "45000000000", OKTMO: "45000000000"},
			},
		},
		{
			name:   "успешно ЮЛ",
			number: "1653001805",
			want: INN{
				Number:        "1653001805",
				EntityType:    INNJuridical,
				AuthorityCode: "1653",
				Region:        dict.Region{Code: "16", Title: "Республика Татарстан (Татарстан)", OKATO: "92000000000", OKTMO: "92000000000"},
			},
		},
		{
			name:    "ошибка контрольной сумма ФЛ 1 разряд",
			number:  "772170926470",
			wantErr: ErrCheckSum,
		},
		{
			name:    "ошибка контрольной сумма ФЛ 2 разряд",
			number:  "772170926461",
			wantErr: ErrCheckSum,
		},
		{
			name:    "ошибка контрольной сумма ЮЛ",
			number:  "1653001806",
			wantErr: ErrCheckSum,
		},
		{
			name:    "ошибка длина",
			number:  "165",
			wantErr: ErrLen,
		},
		{
			name:    "ошибка символ ЮЛ",
			number:  "1653001xxx",
			wantErr: ErrChar,
		},
		{
			name:    "ошибка символ ФЛ",
			number:  "7721709264хх",
			wantErr: ErrChar,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseINN(tt.number)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ParseINN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseINN() got = %v, want %v", got, tt.want)
			}
		})
	}
}
