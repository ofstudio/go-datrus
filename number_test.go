package datrus

import (
	"errors"
	"testing"
)

func Test_validateNumber(t *testing.T) {
	type args struct {
		number      string
		expectedLen int
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "успешно фикс длина",
			args: args{number: "123456", expectedLen: 6},
		},
		{
			name: "успешно произвольная длина",
			args: args{number: "123456", expectedLen: lenAny},
		},
		{
			name:    "недопустимый символ фикс длина",
			args:    args{number: "12*456", expectedLen: 6},
			wantErr: ErrChar,
		},
		{
			name:    "недопустимый символ произвольная длина",
			args:    args{number: "1**6", expectedLen: lenAny},
			wantErr: ErrChar,
		},
		{
			name:    "неверная длина меньше",
			args:    args{number: "123", expectedLen: 10},
			wantErr: ErrLen,
		},
		{
			name:    "неверная длина больше",
			args:    args{number: "1234567890", expectedLen: 5},
			wantErr: ErrLen,
		},
		{
			name:    "пустая строка фикс длина",
			args:    args{number: "", expectedLen: 5},
			wantErr: ErrLen,
		},
		{
			name:    "пустая строка произвольная длина",
			args:    args{number: "", expectedLen: lenAny},
			wantErr: ErrLen,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateNumber(tt.args.number, tt.args.expectedLen); !errors.Is(err, tt.wantErr) {
				t.Errorf("validateNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
