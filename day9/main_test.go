// Advent of Code 2020, Day 9
package main

import (
	"testing"
)

func Test_findFirstInvalid(t *testing.T) {
	type args struct {
		numbers []int
		window  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{[]int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}, 5}, 127},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstInvalid(tt.args.numbers, tt.args.window); got != tt.want {
				t.Errorf("findFirstInvalid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPart2(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{[]int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}, 127}, 62},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPart2(tt.args.numbers, tt.args.target); got != tt.want {
				t.Errorf("findPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
