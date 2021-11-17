// Advent of Code 2020, Day 1
package main

import (
	"fmt"

	"github.com/ghonzo/advent2020/common"
)

// Day 1: Report Repair
// Part 1 answer: 539851
// Part 2 answer: 212481360
func main() {
	fmt.Println("Advent of Code 2020, Day 1")
	entries := common.ReadIntsFromFile("input.txt")
	fmt.Printf("Part 1: Entries multiplied =  %d\n", part1(entries, 2020))
	fmt.Printf("Part 2: Entries multiplied =  %d\n", part2(entries, 2020))
}

func part1(entries []int, sum int) int {
	// Find the first two numbers that add up to sum, then multiple them
	for i, e := range entries {
		for _, e2 := range entries[i+1:] {
			if e+e2 == sum {
				return e * e2
			}
		}
	}
	panic("Not found")
}

func part2(entries []int, sum int) int {
	// Find the first three numbers that add up to sum, then multiple them
	for i, e := range entries {
		for j, e2 := range entries[i+1:] {
			for _, e3 := range entries[j+1:] {
				if e+e2+e3 == sum {
					return e * e2 * e3
				}
			}
		}
	}
	panic("Not found")
}
