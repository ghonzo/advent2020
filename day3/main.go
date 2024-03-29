// Advent of Code 2020, Day 3
package main

import (
	"fmt"

	"github.com/ghonzo/advent2020/common"
)

// Day 3: Toboggan Trajectory
// Part 1 answer: 294
// Part 2 answer: 5774564250
func main() {
	fmt.Println("Advent of Code 2020, Day 3")
	g := geology{common.ReadStringsFromFile("input.txt")}
	fmt.Printf("Part 1: hit trees =  %d\n", g.hitTreesForSlope(3, 1))
	product := g.hitTreesForSlope(1, 1) * g.hitTreesForSlope(3, 1) * g.hitTreesForSlope(5, 1) * g.hitTreesForSlope(7, 1) * g.hitTreesForSlope(1, 2)
	fmt.Printf("Part 2: hit trees product =  %d\n", product)
}

type geology struct {
	lines []string
}

func (g geology) sizeX() int {
	return len(g.lines[0])
}

func (g geology) sizeY() int {
	return len(g.lines)
}

func (g geology) hitTree(x, y int) bool {
	return g.lines[y][x%g.sizeX()] == '#'
}

func (g geology) hitTreesForSlope(xInc, yInc int) int {
	var x int
	var hitTrees int
	for y := 0; y < g.sizeY(); y += yInc {
		if g.hitTree(x, y) {
			hitTrees++
		}
		x += xInc
	}
	return hitTrees
}
