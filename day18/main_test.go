// Advent of Code 2020, Day 18
package main

import (
	"testing"
)

func Test_evalLine(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"1 + 2 * 3 + 4 * 5 + 6"}, 71},
		{"2", args{"1 + (2 * 3) + (4 * (5 + 6))"}, 51},
		{"3", args{"2 * 3 + (4 * 5)"}, 26},
		{"4", args{"5 + (8 * 3 + 9 + 3 * 4 * 3)"}, 437},
		{"5", args{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}, 12240},
		{"6", args{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}, 13632},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalLine(tt.args.s); got != tt.want {
				t.Errorf("evalLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evalLinePart2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"1 + 2 * 3 + 4 * 5 + 6"}, 231},
		{"2", args{"1 + (2 * 3) + (4 * (5 + 6))"}, 51},
		{"3", args{"2 * 3 + (4 * 5)"}, 46},
		{"4", args{"5 + (8 * 3 + 9 + 3 * 4 * 3)"}, 1445},
		{"5", args{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}, 669060},
		{"6", args{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}, 23340},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalLinePart2(tt.args.s); got != tt.want {
				t.Errorf("evalLinePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
