// Advent of Code 2020, Day 16
package main

import (
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	rules, _, nearbyTickets := readInput(strings.NewReader(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`))
	type args struct {
		rules         []rule
		nearbyTickets []ticket
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{rules, nearbyTickets}, 71},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.rules, tt.args.nearbyTickets); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRulesInFieldOrder(t *testing.T) {
	rules, _, nearbyTickets := readInput(strings.NewReader(`class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`))
	wantNames := [...]string{"row", "class", "seat"}
	rulesInFieldOrder := findRulesInFieldOrder(rules, nearbyTickets)
	for i, name := range wantNames {
		if rulesInFieldOrder[i].name != name {
			t.Errorf("findRulesInFieldOrder() = %v, want %v", rulesInFieldOrder[i].name, name)
		}
	}
}
