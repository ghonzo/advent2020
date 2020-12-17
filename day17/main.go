// Advent of Code 2020, Day 17
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Day 17: Conway Cubes
// Part 1 answer: 362
// Part 2 answer: 1980
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
	fmt.Printf("Part 2. Answer = %d\n", part2(initialState))
}

type point4d struct {
	x, y, z, w int
}

type state map[point4d]bool

func readState(r io.Reader) state {
	s := make(state)
	input := bufio.NewScanner(r)
	var y int
	for input.Scan() {
		for x, c := range input.Bytes() {
			if c == '#' {
				s[point4d{x, y, 0, 0}] = true
			}
		}
		y++
	}
	return s
}

func part1(s state) int {
	for i := 0; i < 6; i++ {
		s = cycle(s, false)
	}
	return len(s)
}

func part2(s state) int {
	for i := 0; i < 6; i++ {
		s = cycle(s, true)
	}
	return len(s)
}

func cycle(s state, use4d bool) state {
	newState := make(state)
	searchSpace := findSearchSpace(s, use4d)
	for k := range searchSpace {
		v := s[k]
		n := neighbors(k, s, use4d)
		if (v && (n == 2 || n == 3)) || (!v && n == 3) {
			newState[k] = true
		}
	}
	return newState
}

// Expand each spot by 1 in every direction
func findSearchSpace(s state, use4d bool) state {
	// Initially we're going to set each thing to true
	searchSpace := make(state)
	for k := range s {
		searchSpace[k] = true
		for _, p := range allSurroundingPoints(k, use4d) {
			searchSpace[p] = true
		}
	}
	return searchSpace
}

func allSurroundingPoints(p point4d, use4d bool) []point4d {
	var surrounding []point4d
	if use4d {
		for x := p.x - 1; x <= p.x+1; x++ {
			for y := p.y - 1; y <= p.y+1; y++ {
				for z := p.z - 1; z <= p.z+1; z++ {
					for w := p.w - 1; w <= p.w+1; w++ {
						if !(x == p.x && y == p.y && z == p.z && w == p.w) {
							surrounding = append(surrounding, point4d{x, y, z, w})
						}
					}
				}
			}
		}
	} else {
		for x := p.x - 1; x <= p.x+1; x++ {
			for y := p.y - 1; y <= p.y+1; y++ {
				for z := p.z - 1; z <= p.z+1; z++ {
					if !(x == p.x && y == p.y && z == p.z) {
						surrounding = append(surrounding, point4d{x, y, z, 0})
					}
				}
			}
		}
	}
	return surrounding
}

func neighbors(p point4d, s state, use4d bool) int {
	var n int
	for _, neighbor := range allSurroundingPoints(p, use4d) {
		if s[neighbor] {
			n++
		}
	}
	return n
}
