// Advent of Code 2020, Day 1
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Day 1: Report Repair
// Part 1 answer: 539851
// Part 2 answer: 212481360
func main() {
	fmt.Println("Advent of Code 2020, Day 1")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	entries := readInts(input)
	fmt.Printf("Part 1: Entries multiplied =  %d\n", part1(entries, 2020))
	fmt.Printf("Part 2: Entries multiplied =  %d\n", part2(entries, 2020))
}

func readInts(r io.Reader) []int {
	var ints []int
	input := bufio.NewScanner(r)
	for input.Scan() {
		i, err := strconv.Atoi(input.Text())
		if err != nil {
			fmt.Printf("Bad input: %s\n", input.Text())
		}
		ints = append(ints, i)
	}
	return ints
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
