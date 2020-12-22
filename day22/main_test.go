// Advent of Code 2020, Day 22
package main

import (
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		arg [2]deck
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{readDecks(strings.NewReader(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`))}, 306},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.arg); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		arg [2]deck
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{readDecks(strings.NewReader(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`))}, 291},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.arg); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
