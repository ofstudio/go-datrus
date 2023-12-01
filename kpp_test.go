package datrus

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/ofstudio/go-datrus/dict"
)

func ExampleParseKPP() {
	kpp, _ := ParseKPP("630343001")
	fmt.Println(kpp.Reason.Code, kpp.Reason.Title)

	// Output: 43 Постановка на учет российской организации по месту нахождения ее филиала

}

func ExampleParseKPP_error() {
	_, err := ParseKPP("630-343-0")
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, ErrKPP), errors.Is(err, ErrChar))
	}

	// Output: ошибка КПП: недопустимый символ
	// true true

}
func TestParseKPP(t *testing.T) {

	tests := []struct {
		name    string
		number  string
		want    KPP
		wantErr error
	}{
		{
			name:   "успешно",
			number: "710543001",
			want: KPP{
				Number:        "710543001",
				Region:        dict.Region{Code: "71", Title: "Тульская область", OKATO: "70000000000", OKTMO: "70000000000"},
				AuthorityCode: "7105",
				Reason: dict.KPPReason{
					Code:  "43",
					Title: "Постановка на учет российской организации по месту нахождения ее филиала",
				},
				SerialNumber: "001",
			},
		},
		{
			name:    "неверная длина",
			number:  "123",
			wantErr: ErrLen,
		},
		{
			name:    "недопустимый символ",
			number:  "12345678*",
			wantErr: ErrChar,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseKPP(tt.number)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ParseKPP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseKPP() got = %v, want %v", got, tt.want)
			}
		})
	}
}
