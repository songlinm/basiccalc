package calc

import (
	"io"
	"strings"
	"testing"
)

func TestPrefix(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantRet float64
		wantErr bool
	}{
		{`basic`, args{strings.NewReader(`+ 1 2`)}, 3.0, false},
		{`single`, args{strings.NewReader(`16`)}, 16.0, false},
		{`arg2_is_op`, args{strings.NewReader(`+ 1 * 2 3`)}, 7, false},
		{`arg1_is_op`, args{strings.NewReader(`+ * 1 2 3`)}, 5, false},
		{`complex`, args{strings.NewReader(`- / 10 + 1 1 * 1 2`)}, 3, false},
		{`complex2`, args{strings.NewReader(`- / * 2 * 5 + 3 6 5 2`)}, 16, false},
		{`minus`, args{strings.NewReader(`- 0 3`)}, -3, false},
		{`divide`, args{strings.NewReader(`/ 3 2`)}, 1.5, false},
		{`divide_by_zero`, args{strings.NewReader(`/ 3 0`)}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := Prefix(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Prefix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRet != tt.wantRet {
				t.Errorf("Prefix() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
