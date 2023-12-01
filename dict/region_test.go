package dict

import (
	"fmt"
	"testing"
)

func ExampleRegionFindByCode() {
	fmt.Println(RegionFindByCode("64").Title)
	fmt.Println(RegionFindByCode("50").Title)
	fmt.Println(RegionFindByCode("00").Title)

	// Output: Саратовская область
	// Московская область
	// Неизвестный регион
}

func ExampleRegionFindByOKATO() {
	fmt.Println(RegionFindByOKATO("64").Title)
	fmt.Println(RegionFindByOKATO("71140").Title)
	fmt.Println(RegionFindByOKATO("25000000000").Title)
	fmt.Println(RegionFindByOKATO("00").Title)

	// Output: Сахалинская область
	// Ямало-Ненецкий автономный округ
	// Иркутская область
	// Неизвестный регион
}

func ExampleRegionFindByOKTMO() {
	fmt.Println(RegionFindByOKTMO("52").Title)
	fmt.Println(RegionFindByOKTMO("71900").Title)
	fmt.Println(RegionFindByOKTMO("83000000000").Title)
	fmt.Println(RegionFindByOKTMO("00").Title)

	// Output: Омская область
	// Ямало-Ненецкий автономный округ
	// Кабардино-Балкарская Республика
	// Неизвестный регион
}

func TestRegionFindByCode(t *testing.T) {
	tests := []struct {
		name string
		code string
		want string
	}{
		{name: "успешно", code: "70", want: "Томская область"},
		{name: "неизвестный", code: "00", want: "Неизвестный регион"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegionFindByCode(tt.code); got.Title != tt.want {
				t.Errorf("RegionFindByCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegionFindByOKATO(t *testing.T) {
	tests := []struct {
		name string
		code string
		want string
	}{
		{name: "успешно 2 цифры", code: "92", want: "Республика Татарстан (Татарстан)"},
		{name: "успешно 11 цифр", code: "03000000000", want: "Краснодарский край"},
		{name: "успешно 5 цифр", code: "71140", want: "Ямало-Ненецкий автономный округ"},
		{name: "успешно 7 цифр", code: "1110000", want: "Ненецкий автономный округ"},
		{name: "не успешно 2 цифры", code: "00", want: RegionUnknown},
		{name: "не успешно 11 цифр", code: "03000123000", want: RegionUnknown},
		{name: "не успешно 5 цифр", code: "03001", want: RegionUnknown}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegionFindByOKATO(tt.code); got.Title != tt.want {
				t.Errorf("RegionFindByOKATO() = %v, want %v", got.Title, tt.want)
			}
		})
	}
}

func TestRegionFindByOKTMO(t *testing.T) {
	tests := []struct {
		name string
		code string
		want string
	}{
		{name: "успешно 2 цифры", code: "92", want: "Республика Татарстан (Татарстан)"},
		{name: "успешно 11 цифр", code: "03000000000", want: "Краснодарский край"},
		{name: "успешно 5 цифр", code: "71900", want: "Ямало-Ненецкий автономный округ"},
		{name: "успешно 7 цифр", code: "1180000", want: "Ненецкий автономный округ"},
		{name: "не успешно 2 цифры", code: "00", want: RegionUnknown},
		{name: "не успешно 11 цифр", code: "03000123000", want: RegionUnknown},
		{name: "не успешно 5 цифр", code: "03001", want: RegionUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegionFindByOKTMO(tt.code); got.Title != tt.want {
				t.Errorf("RegionFindByOKTMO() = %v, want %v", got.Title, tt.want)
			}
		})
	}
}
