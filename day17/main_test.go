// Advent of Code 2020, Day 17
package main

import (
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		s state
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{readState(strings.NewReader(".#.\n..#\n###"))}, 112},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.s); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		s state
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{readState(strings.NewReader(".#.\n..#\n###"))}, 848},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.s); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
