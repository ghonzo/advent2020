// Advent of Code 2020, Day 21
package main

import (
	"strings"
	"testing"
)

func Test_part1And2(t *testing.T) {
	type args struct {
		foods []food
	}
	tests := []struct {
		name  string
		args  args
		want1 int
		want2 string
	}{
		{"example", args{readFoods(strings.NewReader(`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`))}, 5, "mxmxvkd,sqjhc,fvjkl"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1, got2 := part1And2(tt.args.foods); got1 != tt.want1 {
				t.Errorf("part1() = %v, want %v", got1, tt.want1)
			} else if got2 != tt.want2 {
				t.Errorf("part2() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
