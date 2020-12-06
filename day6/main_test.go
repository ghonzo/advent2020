// Advent of Code 2020, Day 6
package main

import (
	"testing"
)

func Test_group_union(t *testing.T) {
	type fields struct {
		person []string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"1", fields{[]string{"abc"}}, 3},
		{"2", fields{[]string{"a", "b", "c"}}, 3},
		{"3", fields{[]string{"ab", "ac"}}, 3},
		{"4", fields{[]string{"a", "a", "a", "a"}}, 1},
		{"5", fields{[]string{"b"}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := group{
				person: tt.fields.person,
			}
			if got := g.union(); got != tt.want {
				t.Errorf("group.union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_group_intersection(t *testing.T) {
	type fields struct {
		person []string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"1", fields{[]string{"abc"}}, 3},
		{"2", fields{[]string{"a", "b", "c"}}, 0},
		{"3", fields{[]string{"ab", "ac"}}, 1},
		{"4", fields{[]string{"a", "a", "a", "a"}}, 1},
		{"5", fields{[]string{"b"}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := group{
				person: tt.fields.person,
			}
			if got := g.intersection(); got != tt.want {
				t.Errorf("group.intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
