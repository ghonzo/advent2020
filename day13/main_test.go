// Advent of Code 2020, Day 13
package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		earliest int
		ids      []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{939, []int{7, 13, 0, 0, 59, 0, 31, 19}}, 295},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.earliest, tt.args.ids); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		ids []int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"1", args{[]int{7, 13, 0, 0, 59, 0, 31, 19}}, 1068781},
		{"2", args{[]int{17, 0, 13, 19}}, 3417},
		{"3", args{[]int{67, 7, 59, 61}}, 754018},
		{"4", args{[]int{67, 0, 7, 59, 61}}, 779210},
		{"5", args{[]int{67, 7, 0, 59, 61}}, 1261476},
		{"6", args{[]int{1789, 37, 47, 1889}}, 1202161486},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.ids); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
