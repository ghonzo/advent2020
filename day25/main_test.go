// Advent of Code 2020, Day 25
package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		cpk int
		dpk int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{5764801, 17807724}, 14897079},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.cpk, tt.args.dpk); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
