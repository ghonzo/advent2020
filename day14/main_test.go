// Advent of Code 2020, Day 14
package main

import (
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		instructions []string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"example", args{readInput(strings.NewReader(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`))}, 165},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.instructions); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		instructions []string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"example", args{readInput(strings.NewReader(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`))}, 208},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.instructions); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
