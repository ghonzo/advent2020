// Advent of Code 2020, Day 25
package main

import (
	"fmt"
)

// Day 25: Combo Breaker
// Part 1 answer: 1478097
// Part 2 answer: THERE IS NO PART 2!
func main() {
	fmt.Println("Advent of Code 2020, Day 25")
	cardPublicKey := 9232416
	doorPublicKey := 14144084
	fmt.Printf("Part 1. Answer = %d\n", part1(cardPublicKey, doorPublicKey))
}

func part1(cpk, dpk int) int {
	card := 7
	for {
		card = transform(card)
		dpk = transform(dpk)
		if card == cpk {
			return dpk
		}
	}
}

func transform(n int) int {
	return n * n % 20201227
}
