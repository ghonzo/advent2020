// Advent of Code 2020, Day 11
package main

import (
	"strings"
	"testing"

	"github.com/ghonzo/advent2020/common"
)

func Test_part1(t *testing.T) {
	type args struct {
		seatMap common.Grid
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadArraysGrid(strings.NewReader(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`))}, 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.seatMap); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		seatMap common.Grid
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadArraysGrid(strings.NewReader(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`))}, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.seatMap); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
