// Advent of Code 2020, Day 16
package main

import (
	"fmt"
)

// Day 15: Rambunctious Recitation
// Part 1 answer: 249
// Part 2 answer: 41687
func main() {
	fmt.Println("Advent of Code 2020, Day 15")
	input := []int{15, 12, 0, 14, 3, 1}
	fmt.Printf("Part 1. Answer = %d\n", part1(input))
	fmt.Printf("Part 2. Answer = %d\n", part2(input))
}

// I could have ditched this and just used the algo for part2, but I kept
// this here for posterity
func part1(input []int) int {
	spoken := input
	for turn := len(input); turn < 2020; turn++ {
		spoken = append(spoken, nextNumber(spoken))
	}
	return spoken[2019]
}

func nextNumber(spoken []int) int {
	lastSpoken := spoken[len(spoken)-1]
	for i := len(spoken) - 2; i >= 0; i-- {
		if spoken[i] == lastSpoken {
			return len(spoken) - 1 - i
		}
	}
	return 0
}

func part2(input []int) int {
	// Number to last turn it was uttered
	spokenMap := make(map[int]int)
	for turn, number := range input[:len(input)-1] {
		spokenMap[number] = turn
	}
	lastSpoken := input[len(input)-1]
	for turn := len(input); turn < 30000000; turn++ {
		v, ok := spokenMap[lastSpoken]
		spokenMap[lastSpoken] = turn - 1
		if !ok {
			lastSpoken = 0
		} else {
			lastSpoken = turn - 1 - v
		}
	}
	return lastSpoken
}
