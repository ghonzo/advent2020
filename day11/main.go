// Advent of Code 2020, Day 11
package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/ghonzo/advent2020/common"
)

// Day 11: Seating System
// Part 1 answer: 2310
// Part 2 answer: 2074
func main() {
	fmt.Println("Advent of Code 2020, Day 11")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	seatMap := common.ReadArraysGrid(input)
	fmt.Printf("Part 1. Answer = %d\n", part1(seatMap))
	fmt.Printf("Part 2. Answer = %d\n", part2(seatMap))
}

const (
	occupied = '#'
	empty    = 'L'
	floor    = '.'
)

func part1(seatMap common.Grid) int {
	for i := 0; ; i++ {
		newSeatMap := seatMap.Clone()
		for pt := range seatMap.AllPoints() {
			newSeatMap.Set(pt, applyPart1Rule(seatMap, pt))
		}
		// let's use DeepEqual ... until we can't
		if reflect.DeepEqual(seatMap, newSeatMap) {
			return seatMap.Count(occupied)
		}
		seatMap = newSeatMap
	}
}

func applyPart1Rule(seatMap common.Grid, pt common.Point) byte {
	current := seatMap.Get(pt)
	if current != floor {
		switch n := countSurroundingOccupied(seatMap, pt); current {
		case occupied:
			if n >= 4 {
				return empty
			}
		case empty:
			if n == 0 {
				return occupied
			}
		}
	}
	// No change
	return current
}

func countSurroundingOccupied(seatMap common.Grid, center common.Point) int {
	var count int
	for _, offset := range common.AllDirections {
		if v, _ := seatMap.CheckedGet(center.Add(offset)); v == occupied {
			count++
		}
	}
	return count
}

func part2(seatMap common.Grid) int {
	for i := 0; ; i++ {
		newSeatMap := seatMap.Clone()
		for pt := range seatMap.AllPoints() {
			newSeatMap.Set(pt, applyPart2Rule(seatMap, pt))
		}
		// let's use DeepEqual ... until we can't
		if reflect.DeepEqual(seatMap, newSeatMap) {
			return seatMap.Count(occupied)
		}
		seatMap = newSeatMap
	}
}

func applyPart2Rule(seatMap common.Grid, pt common.Point) byte {
	current := seatMap.Get(pt)
	if current != floor {
		switch n := countRadialOccupied(seatMap, pt); current {
		case occupied:
			if n >= 5 {
				return empty
			}
		case empty:
			if n == 0 {
				return occupied
			}
		}
	}
	// No change
	return current
}

func countRadialOccupied(seatMap common.Grid, center common.Point) int {
	var count int
	for _, direction := range common.AllDirections {
		if seeOccupied(seatMap, center, direction) {
			count++
		}
	}
	return count
}

func seeOccupied(seatMap common.Grid, pt common.Point, direction common.Point) bool {
	for {
		pt = pt.Add(direction)
		if v, ok := seatMap.CheckedGet(pt); !ok {
			// walked off the edge
			return false
		} else if v == occupied {
			return true
		} else if v == empty {
			return false
		}
		// must be floor ... keep looking
	}
}
