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

// Add returns a new Point with x- and y-coordinates added
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

// AllDirections is a slice of all the different offsets that repesent the different
// directions around a point. Does not include "center" or "zero"
var AllDirections = []Point{UL, U, UR, L, R, DL, D, DR}

// Grid represents a mutable 2D rectangle with a value at each integer coordinate
type Grid interface {
	Size() Point
	Get(coord Point) byte
	CheckedGet(coord Point) (v byte, ok bool)
	Set(coord Point, b byte)
	Count(b byte) int
	AllPoints() <-chan Point
	Clone() Grid
}

// ArraysGrid is a Grid that has an underlying representation of [][]byte
type ArraysGrid [][]byte

// Size returns a Point that represents the dimensions of the Grid
func (g *ArraysGrid) Size() Point {
	return Point{len((*g)[0]), len(*g)}
}

// Get returns the value at the given coordinate
//
// If the coordinate is outside the dimensions of the Grid, this will throw an error. You may
// want to use CheckedGet instead
func (g *ArraysGrid) Get(coord Point) byte {
	return (*g)[coord.y][coord.x]
}

// CheckedGet returns the value at the given coordinate (if present), as well as an ok value.
//
// If the coordinate is present, it will return the value and an ok value of true. If it is not,
// it will return 0 and an ok value of false.
func (g *ArraysGrid) CheckedGet(coord Point) (byte, bool) {
	size := g.Size()
	if coord.x < 0 || coord.x >= size.x || coord.y < 0 || coord.y >= size.y {
		return 0, false
	}
	return g.Get(coord), true
}

// Set sets a value at the given coordinate.
//
// If the coordinate is outside the bounds, this will throw an error.
func (g *ArraysGrid) Set(coord Point, b byte) {
	(*g)[coord.y][coord.x] = b
}

// Count the number of points in the Grid that have the given value.
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

// Clone returns a copy of the Grid, leaving the original untouched.
func (g *ArraysGrid) Clone() Grid {
	size := g.Size()
	clone := make(ArraysGrid, size.y)
	for row := range *g {
		clone[row] = make([]byte, size.x)
		copy(clone[row], (*g)[row])
	}
	return &clone
}

// AllPoints returns a channel of all the points in the Grid.
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

// ReadArraysGrid parses the lines and created a Grid with the y-dimension given by the number of lines
// and the x-dimension given by the length of the first line.
func ReadArraysGrid(r io.Reader) *ArraysGrid {
	var grid ArraysGrid
	input := bufio.NewScanner(r)
	for input.Scan() {
		grid = append(grid, []byte(input.Text()))
	}
	return &grid
}
