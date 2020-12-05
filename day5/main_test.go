// Advent of Code 2020, Day 5
package main

import (
	"testing"
)

func Test_boardingPass_row(t *testing.T) {
	tests := []struct {
		name string
		b    boardingPass
		want int
	}{
		{"1", boardingPass("FBFBBFFRLR"), 44},
		{"2", boardingPass("BFFFBBFRRR"), 70},
		{"3", boardingPass("FFFBBBFRRR"), 14},
		{"4", boardingPass("BBFFBBFRLL"), 102},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.row(); got != tt.want {
				t.Errorf("boardingPass.row() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardingPass_column(t *testing.T) {
	tests := []struct {
		name string
		b    boardingPass
		want int
	}{
		{"1", boardingPass("FBFBBFFRLR"), 5},
		{"2", boardingPass("BFFFBBFRRR"), 7},
		{"3", boardingPass("FFFBBBFRRR"), 7},
		{"4", boardingPass("BBFFBBFRLL"), 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.column(); got != tt.want {
				t.Errorf("boardingPass.column() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardingPass_seatID(t *testing.T) {
	tests := []struct {
		name string
		b    boardingPass
		want int
	}{
		{"1", boardingPass("FBFBBFFRLR"), 357},
		{"2", boardingPass("BFFFBBFRRR"), 567},
		{"3", boardingPass("FFFBBBFRRR"), 119},
		{"4", boardingPass("BBFFBBFRLL"), 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.seatID(); got != tt.want {
				t.Errorf("boardingPass.seatID() = %v, want %v", got, tt.want)
			}
		})
	}
}
