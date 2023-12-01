package datrus

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleParseSNILS() {
	snils, _ := ParseSNILS("016 480 666 49")
	fmt.Println(snils)

	snils, _ = ParseSNILS("477-004-761 90")
	fmt.Println(snils)

	snils, _ = ParseSNILS("621 248 259 65")
	fmt.Println(snils)

	// Output: 01648066649
	// 47700476190
	// 62124825965

}

func ExampleParseSNILS_error() {
	_, err := ParseSNILS("111 111 111 00")
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, ErrSNILS), errors.Is(err, ErrCheckSum))
	}

	_, err = ParseSNILS("111 abc def 00")
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, ErrSNILS), errors.Is(err, ErrLen))
	}

	// Output: ошибка SNILS: ошибка контрольной суммы
	// true true
	// ошибка SNILS: неверная длина
	// true true

}

func ExampleSNILS_Format() {
	snils := MustParseSNILS("715 398 174 20")

	fmt.Println(snils)
	fmt.Println(snils.Format(SNILSDashSpace))
	fmt.Println(snils.Format(SNILSDash))
	fmt.Println(snils.Format(SNILSSpace))

	// Output: 71539817420
	// 715-398-174 20
	// 715-398-174-20
	// 715 398 174 20

}

func TestParseSNILS(t *testing.T) {
	tests := []struct {
		name    string
		number  string
		want    SNILS
		wantErr error
	}{
		{name: "успешно пробелы", number: "276 488 905 42", want: "27648890542"},
		{name: "успешно дефисы", number: "002-064-585-96", want: "00206458596"},
		{name: "успешно дефисы и пробел", number: "774-112-048 81", want: "77411204881"},
		{name: "успешно слитно", number: "77241387515", want: "77241387515"},
		{name: "ошибка формат", number: "77 241 387 515", wantErr: ErrLen},
		{name: "ошибка длина", number: "908 035 685", wantErr: ErrLen},
		{name: "ошибка символы", number: "589*088*281*65", wantErr: ErrLen},
		{name: "ошибка контрольная сумма", number: "200 746 095 00", wantErr: ErrCheckSum},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSNILS(tt.number)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ParseSNILS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseSNILS() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRandomSNILS(t *testing.T) {
	for i := 0; i < 10_000; i++ {
		snils, err := NewRandomSNILS()

		// check no error
		if err != nil {
			t.Errorf("NewRandomSNILS() error = %v", err)
			return
		}

		// check snils
		if _, err = ParseSNILS(string(snils)); err != nil {
			t.Errorf("NewRandomSNILS() error = %v", err)
			return
		}
	}
}
