// Advent of Code 2020, Day 17
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Day 17: Conway Cubes
// Part 1 answer:
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2020, Day 17")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	initialState := readState(input)
	fmt.Printf("Part 1. Answer = %d\n", part1(initialState))
	//fmt.Printf("Part 2. Answer = %d\n", part2(rules, yourTicket, nearbyTickets))
}

type point3d struct {
	x, y, z int
}

type point4d struct {
	point3d
	w int
}

type state map[point3d]bool

type state4d map[point4d]bool

func readState(r io.Reader) state {
	s := make(state)
	input := bufio.NewScanner(r)
	var y int
	for input.Scan() {
		for x, c := range input.Bytes() {
			if c == '#' {
				s[point3d{x, y, 0}] = true
			}
		}
		y++
	}
	return s
}

func part1(s state) int {
	for i := 0; i < 6; i++ {
		s = cycle(s)
	}
	return len(s)
}

func cycle(s state) state {
	newState := make(state)
	searchSpace := findSearchSpace(s)
	for k := range searchSpace {
		v := s[k]
		n := neighbors(k, s)
		if (v && (n == 2 || n == 3)) || (!v && n == 3) {
			newState[k] = true
		}
	}
	return newState
}

// Expand each spot by 1 in every direction
func findSearchSpace(s state) state {
	// Initially we're going to set each thing to true
	searchSpace := make(state)
	for k := range s {
		searchSpace[k] = true
		for _, p := range allSurroundingPoints(k) {
			searchSpace[p] = true
		}
	}
	return searchSpace
}

func allSurroundingPoints(p point3d) []point3d {
	var surrounding []point3d
	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			for z := p.z - 1; z <= p.z+1; z++ {
				if !(x == p.x && y == p.y && z == p.z) {
					surrounding = append(surrounding, point3d{x, y, z})
				}
			}
		}
	}
	return surrounding
}

func neighbors(p point3d, s state) int {
	var n int
	for _, neighbor := range allSurroundingPoints(p) {
		if s[neighbor] {
			n++
		}
	}
	return n
}
