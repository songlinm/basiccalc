package calc

import (
	"io"
	"strings"
	"testing"
)

func TestInfix(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantRet float64
		wantErr bool
	}{
		{`one_value`, args{strings.NewReader(`2048`)}, 2048, false},
		{`no_parenthesis`, args{strings.NewReader(`7 + 99`)}, 106, false},
		{`simple`, args{strings.NewReader(`( 1 + 2 )`)}, 3, false},
		{`arg1_parenthesis`, args{strings.NewReader(`( ( 1 * 2 ) + 3 )`)}, 5, false},
		{`arg2_parenthesis`, args{strings.NewReader(`( 1 + ( 2 * 3 ) )`)}, 7, false},
		{`complex`, args{strings.NewReader(`( ( ( 1 + 1 ) / 10 ) - ( 1 * 2 ) )`)}, -1.8, false},
		{`complex2`, args{strings.NewReader(`( ( 1 + 9 ) / 2 ) * ( 3 + 8 )`)}, 55, false},
		{`complex3`, args{strings.NewReader(`( ( 8 + ( 1 + 9 ) ) / 2 ) * ( 3 + 8 )`)}, 99, false},
		{`divide_by_zero`, args{strings.NewReader(`( ( ( 1 + 1 ) / 0 ) - ( 1 * 2 ) )`)}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := Infix(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Infix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRet != tt.wantRet {
				t.Errorf("Infix() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
