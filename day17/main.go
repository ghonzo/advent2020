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
		s = cycle(s, allSurroundingPoints3d)
	}
	return len(s)
}

func part2(s state) int {
	for i := 0; i < 6; i++ {
		s = cycle(s, allSurroundingPoints4d)
	}
	return len(s)
}

type surroundingPointsFunc func(point4d) []point4d

func cycle(s state, f surroundingPointsFunc) state {
	newState := make(state)
	searchSpace := findSearchSpace(s, f)
	// Iterate over all of the points in the search space
	for p := range searchSpace {
		active := s[p]
		n := neighbors(p, s, f)
		if (active && (n == 2 || n == 3)) || (!active && n == 3) {
			newState[p] = true
		}
	}
	return newState
}

func findSearchSpace(s state, f surroundingPointsFunc) state {
	// We are going to return a state with every new point that needs to be checked
	searchSpace := make(state)
	// Iterate over all the existing active points
	for p := range s {
		searchSpace[p] = true
		// And also expand to all surrounding points
		for _, surroundingPoint := range f(p) {
			searchSpace[surroundingPoint] = true
		}
	}
	return searchSpace
}

// Don't delve into the w dimension
func allSurroundingPoints3d(p point4d) []point4d {
	var surrounding []point4d
	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			for z := p.z - 1; z <= p.z+1; z++ {
				if !(x == p.x && y == p.y && z == p.z) {
					surrounding = append(surrounding, point4d{x, y, z, 0})
				}
			}
		}
	}
	return surrounding
}

func allSurroundingPoints4d(p point4d) []point4d {
	var surrounding []point4d
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
	return surrounding
}

func neighbors(p point4d, s state, f surroundingPointsFunc) int {
	var n int
	for _, neighbor := range f(p) {
		if s[neighbor] {
			n++
		}
	}
	return n
}
