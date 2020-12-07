// Advent of Code 2020, Day 7
package main

import (
	"strings"
	"testing"
)

func Test_countContainedIn(t *testing.T) {
	type args struct {
		target color
		rules  map[color][]colorAndQty
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{color("shiny gold"), readRules(strings.NewReader(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`))}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countContainedIn(tt.args.target, tt.args.rules); got != tt.want {
				t.Errorf("countContainedIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countContained(t *testing.T) {
	type args struct {
		bag   color
		rules map[color][]colorAndQty
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{color("shiny gold"), readRules(strings.NewReader(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`))}, 33},
		{"example2", args{color("shiny gold"), readRules(strings.NewReader(`shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`))}, 127},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countContained(tt.args.bag, tt.args.rules); got != tt.want {
				t.Errorf("countContained() = %v, want %v", got, tt.want)
			}
		})
	}
}
