// Advent of Code 2020, Day 6
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Day 6: Custom Customs
// Part 1 answer: 6903
// Part 2 answer: 3493
func main() {
	fmt.Println("Advent of Code 2020, Day 6")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	groups := readGroups(input)
	var part1Count, part2Count int
	for _, g := range groups {
		part1Count += g.union()
		part2Count += g.intersection()
	}
	fmt.Printf("Part 1: count =  %d\n", part1Count)
	fmt.Printf("Part 2: count =  %d\n", part2Count)
}

type group struct {
	person []string
}

func readGroups(r io.Reader) []group {
	var groups []group
	input := bufio.NewScanner(r)
	var g group
	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			groups = append(groups, g)
			g = group{}
		} else {
			g.person = append(g.person, line)
		}
	}
	groups = append(groups, g)
	return groups
}

func (g group) union() int {
	questions := make(map[rune]bool)
	for _, p := range g.person {
		for _, q := range p {
			questions[q] = true
		}
	}
	return len(questions)
}

func (g group) intersection() int {
	var questions map[rune]bool
	for _, p := range g.person {
		personQuestions := make(map[rune]bool)
		for _, q := range p {
			if questions == nil || questions[q] {
				personQuestions[q] = true
			}
		}
		questions = personQuestions
	}
	return len(questions)
}
