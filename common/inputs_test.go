package common

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestReadInts(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"normal", args{strings.NewReader("0\n1\n+2\n-3")}, []int{0, 1, 2, -3}},
		{"empty", args{strings.NewReader("")}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadInts(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadInts_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ReadInts should have panic'd with a non-integer value")
		}
	}()

	ReadInts(strings.NewReader("0\nNot an int\n2"))
}
