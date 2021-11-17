// Advent of Code 2020, Day 9
package main

import (
	"fmt"

	"github.com/ghonzo/advent2020/common"
)

// Day 9: Encoding Error
// Part 1 answer: 675280050
// Part 2 answer: 96081673
func main() {
	fmt.Println("Advent of Code 2020, Day 9")
	numbers := common.ReadIntsFromFile("input.txt")
	invalid := findFirstInvalid(numbers, 25)
	fmt.Printf("Part 1. answer = %d\n", invalid)
	part2 := findPart2(numbers, invalid)
	fmt.Printf("Part 2. answer = %d\n", part2)
}

func findFirstInvalid(numbers []int, window int) int {
	for i := window; i < len(numbers); i++ {
		if !validNumber(numbers, window, i) {
			return numbers[i]
		}
	}
	panic("Not found")
}

func validNumber(numbers []int, window, index int) bool {
	target := numbers[index]
	for a := index - window; a < index-1; a++ {
		for b := a + 1; b < index; b++ {
			if numbers[a] == numbers[b] {
				continue
			}
			if numbers[a]+numbers[b] == target {
				return true
			}
		}
	}
	return false
}

func findPart2(numbers []int, target int) int {
	for start := 0; start < len(numbers); start++ {
		sum := 0
		smallest := target
		largest := 0
		for i := start; i < len(numbers); i++ {
			v := numbers[i]
			if v < smallest {
				smallest = v
			}
			if v > largest {
				largest = v
			}
			sum += v
			if sum == target {
				return smallest + largest
			}
			if sum > target {
				break
			}
		}
	}
	panic("Not found")
}
