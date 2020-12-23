// Advent of Code 2020, Day 23
package main

import (
	"fmt"
	"strconv"
)

// Day 23: Crab Cups
// Part 1 answer: 98645732
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2020, Day 23")
	input := "364289715"
	fmt.Printf("Part 1. Answer = %s\n", part1(input))
	//fmt.Printf("Part 2. Answer = %d\n", part2(decks))
}

type state struct {
	cups            []int
	currentCupIndex int
}

// Remove the three cups following currentCup
func (s *state) removeCups() []int {
	removed := make([]int, 3)
	if s.currentCupIndex+3 < len(s.cups) {
		copy(removed, s.cups[s.currentCupIndex+1:s.currentCupIndex+4])
		s.cups = append(s.cups[:s.currentCupIndex+1], s.cups[s.currentCupIndex+4:]...)
		return removed
	}
	removedFromFront := (s.currentCupIndex + 4) % len(s.cups)
	copy(removed, append(s.cups[s.currentCupIndex+1:], s.cups[:removedFromFront]...))
	s.cups = s.cups[removedFromFront : s.currentCupIndex+1]
	s.currentCupIndex -= removedFromFront
	return removed
}

func (s *state) findDesintationCup() int {
	// What is the current cup value?
	currentCupValue := s.cups[s.currentCupIndex]
	for cupToFind := currentCupValue - 1; cupToFind > 0; cupToFind-- {
		for i, v := range s.cups {
			if v == cupToFind {
				return i
			}
		}
	}
	// Nope, find the highest
	var highest, highestIndex int
	for i, v := range s.cups {
		if v > highest {
			highest = v
			highestIndex = i
		}
	}
	return highestIndex
}

func (s *state) insertCups(index int, c []int) {
	s.cups = append(s.cups[:index+1], append(c, s.cups[index+1:]...)...)
	if s.currentCupIndex > index {
		s.currentCupIndex += len(c)
	}
}

func (s *state) selectNewCurrentCup() {
	s.currentCupIndex = (s.currentCupIndex + 1) % len(s.cups)
}

func (s *state) String() string {
	var str string
	for i, v := range s.cups {
		if i == s.currentCupIndex {
			str += "("
		}
		str += strconv.Itoa(v)
		if i == s.currentCupIndex {
			str += ")"
		}
	}
	return str
}

func part1(input string) string {
	cups := make([]int, 0, len(input))
	for _, ch := range input {
		cups = append(cups, int(ch)-'0')
	}
	s := &state{cups, 0}
	fmt.Println("Initial State ", s)
	for move := 0; move < 100; move++ {
		// Let's find the destination cup
		removedCups := s.removeCups()
		destinationCupIndex := s.findDesintationCup()
		s.insertCups(destinationCupIndex, removedCups)
		s.selectNewCurrentCup()
		fmt.Println(s)
	}
	return s.String()
}
