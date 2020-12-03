// Advent of Code 2020, Day 3
package main

import (
	"strings"
	"testing"
)

func Test_geology_hitTreesForSlope(t *testing.T) {
	type args struct {
		xInc int
		yInc int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3,1", args{3, 1}, 7},
		{"1,1", args{1, 1}, 2},
		{"5,1", args{5, 1}, 3},
		{"7,1", args{7, 1}, 4},
		{"1,2", args{1, 2}, 2},
	}
	g := readGeology(strings.NewReader(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.hitTreesForSlope(tt.args.xInc, tt.args.yInc); got != tt.want {
				t.Errorf("geology.hitTreesForSlope() = %v, want %v", got, tt.want)
			}
		})
	}
}
