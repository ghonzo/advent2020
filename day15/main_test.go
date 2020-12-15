// Advent of Code 2020, Day 15
package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{[]int{0, 3, 6}}, 436},
		{"1", args{[]int{1, 3, 2}}, 1},
		{"2", args{[]int{2, 1, 3}}, 10},
		{"3", args{[]int{1, 2, 3}}, 27},
		{"4", args{[]int{2, 3, 1}}, 78},
		{"5", args{[]int{3, 2, 1}}, 438},
		{"6", args{[]int{3, 1, 2}}, 1836},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{[]int{0, 3, 6}}, 175594},
		{"1", args{[]int{1, 3, 2}}, 2578},
		{"2", args{[]int{2, 1, 3}}, 3544142},
		{"3", args{[]int{1, 2, 3}}, 261214},
		{"4", args{[]int{2, 3, 1}}, 6895259},
		{"5", args{[]int{3, 2, 1}}, 18},
		{"6", args{[]int{3, 1, 2}}, 362},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
