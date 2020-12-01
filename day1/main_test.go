// Advent of Code 2020, Day 1
package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		entries []int
		sum     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample", args{[]int{1721, 979, 366, 299, 675, 1456}, 2020}, 514579},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.entries, tt.args.sum); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		entries []int
		sum     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample", args{[]int{1721, 979, 366, 299, 675, 1456}, 2020}, 241861950},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.entries, tt.args.sum); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
