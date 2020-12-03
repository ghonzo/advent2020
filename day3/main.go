// Advent of Code 2020, Day 3
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Day 3: Toboggan Trajectory
// Part 1 answer: 294
// Part 2 answer: 5774564250
func main() {
	fmt.Println("Advent of Code 2020, Day 3")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	g := readGeology(input)
	fmt.Printf("Part 1: hit trees =  %d\n", g.hitTreesForSlope(3, 1))
	product := g.hitTreesForSlope(1, 1) * g.hitTreesForSlope(3, 1) * g.hitTreesForSlope(5, 1) * g.hitTreesForSlope(7, 1) * g.hitTreesForSlope(1, 2)
	fmt.Printf("Part 2: hit trees product =  %d\n", product)
}

type pos struct {
	x, y int
}

type geology struct {
	trees        map[pos]bool
	sizeX, sizeY int
}

func (g geology) hitTree(x, y int) bool {
	return g.trees[pos{x % g.sizeX, y}]
}

func readGeology(r io.Reader) geology {
	var g geology
	g.trees = make(map[pos]bool)
	input := bufio.NewScanner(r)
	var y int
	for input.Scan() {
		line := input.Text()
		g.sizeX = len(line)
		for x, ch := range line {
			g.trees[pos{x, y}] = (ch == '#')
		}
		y++
	}
	g.sizeY = y
	return g
}

func (g geology) hitTreesForSlope(xInc, yInc int) int {
	var x int
	var hitTrees int
	for y := 0; y < g.sizeY; y += yInc {
		if g.hitTree(x, y) {
			hitTrees++
		}
		x += xInc
	}
	return hitTrees
}
