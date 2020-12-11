// Package common provides common data structures and utility functions for Advent of Code 2020
package common

import (
	"bufio"
	"io"
)

// Point is an immutable data structure representing an X and Y coordinate pair.
//
// I know we could have used image.Point, but I wanted to enforce immutability and plus
// it's just fun to write.
type Point struct {
	x, y int
}

// X returns the x-coordinate
func (p Point) X() int {
	return p.x
}

// Y returns the y-coordinate
func (p Point) Y() int {
	return p.y
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

// All of these directions (Up Down Left Right) assume "UP" and "LEFT" mean -1 while "DOWN" and "RIGHT" mean +1
var (
	UL = Point{-1, -1}
	U  = Point{0, -1}
	UR = Point{1, -1}
	L  = Point{-1, 0}
	/* skip 0,0 */
	R  = Point{1, 0}
	DL = Point{-1, 1}
	D  = Point{0, 1}
	DR = Point{1, 1}
)

var AllDirections = []Point{UL, U, UR, L, R, DL, D, DR}

type Grid interface {
	Size() Point
	Get(coord Point) byte
	CheckedGet(coord Point) (v byte, ok bool)
	Set(coord Point, b byte)
	Count(b byte) int
	AllPoints() <-chan Point
	Clone() Grid
}

type ArraysGrid [][]byte

func (g *ArraysGrid) Size() Point {
	return Point{len((*g)[0]), len(*g)}
}

func (g *ArraysGrid) Get(coord Point) byte {
	return (*g)[coord.y][coord.x]
}

func (g *ArraysGrid) CheckedGet(coord Point) (byte, bool) {
	size := g.Size()
	if coord.x < 0 || coord.x >= size.x || coord.y < 0 || coord.y >= size.y {
		return 0, false
	}
	return g.Get(coord), true
}

func (g *ArraysGrid) Set(coord Point, b byte) {
	(*g)[coord.y][coord.x] = b
}

func (g *ArraysGrid) Count(b byte) int {
	size := g.Size()
	var count int
	for y := 0; y < size.y; y++ {
		for x := 0; x < size.x; x++ {
			if (*g)[y][x] == b {
				count++
			}
		}
	}
	return count
}

func (g *ArraysGrid) Clone() Grid {
	size := g.Size()
	clone := make(ArraysGrid, size.y)
	for row := range *g {
		clone[row] = make([]byte, size.x)
		copy(clone[row], (*g)[row])
	}
	return &clone
}

func (g *ArraysGrid) AllPoints() <-chan Point {
	ch := make(chan Point)
	go func() {
		size := g.Size()
		for y := 0; y < size.y; y++ {
			for x := 0; x < size.x; x++ {
				ch <- Point{x, y}
			}
		}
		close(ch)
	}()
	return ch
}

func ReadArraysGrid(r io.Reader) *ArraysGrid {
	var grid ArraysGrid
	input := bufio.NewScanner(r)
	for input.Scan() {
		grid = append(grid, []byte(input.Text()))
	}
	return &grid
}
