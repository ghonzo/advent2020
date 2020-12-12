// Advent of Code 2020, Day 12
package main

import (
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		ins []instruction
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{readInstructions(strings.NewReader(`F10
N3
F7
R90
F11`))}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.ins); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		ins []instruction
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{readInstructions(strings.NewReader(`F10
N3
F7
R90
F11`))}, 286},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.ins); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
