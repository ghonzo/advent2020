// Advent of Code 2020, Day 10
package main

import (
	"fmt"

	"github.com/ghonzo/advent2020/common"
)

// Day 10: Adapter Array
// Part 1 answer: 2030
// Part 2 answer: 42313823813632
func main() {
	fmt.Println("Advent of Code 2020, Day 10")
	numbers := common.ReadIntsFromFile("input.txt")
	fmt.Printf("Part 1. Answer = %d\n", part1(numbers))
	fmt.Printf("Part 2. Answer = %d\n", part2(numbers))
}

func part1(numbers []int) int {
	adaptors := make(map[int]bool)
	for _, n := range numbers {
		adaptors[n] = true
	}
	var lastJolt, skip1, skip3 int
	for jolt := 1; jolt-lastJolt < 4; jolt++ {
		if adaptors[jolt] {
			if jolt-lastJolt == 1 {
				skip1++
			} else if jolt-lastJolt == 3 {
				skip3++
			}
			lastJolt = jolt
		}
	}
	skip3++
	return skip1 * skip3
}

func part2(numbers []int) int {
	/*
	 * So here's the deal. We can figure out the number of combinations in each
	 * subgroup (separated by a 3-jolt jump), and then multiply them all together.
	 * It turns out that the number of combinations for each subgroup is given by
	 * the Tribonacci sequence, where each term is the sum of the preceding three
	 * terms. So with this array, you find out the number of consecutive 1-jolt
	 * jumps and use that as the index to return the combinatorial number.
	 */
	tribonacci := [...]int{ /* 0, 0, */ 1, 1, 2, 4, 7, 13, 24}
	adaptors := make(map[int]bool)
	for _, n := range numbers {
		adaptors[n] = true
	}
	arrangements := 1
	var lastJolt, skip1Run int
	for jolt := 1; jolt-lastJolt < 4; jolt++ {
		if adaptors[jolt] {
			if jolt-lastJolt == 1 {
				skip1Run++
			} else if jolt-lastJolt == 3 {
				// Okay, we just had a run of skip1Run 1-jolt jumps
				arrangements *= tribonacci[skip1Run]
				skip1Run = 0
			}
			lastJolt = jolt
		}
	}
	// Don't forget to finish it off
	arrangements *= tribonacci[skip1Run]
	return arrangements
}
