// Advent of Code 2020, Day 24
package main

import (
	"fmt"

	"github.com/ghonzo/advent2020/common"
)

// Day 24: Lobby Layout
// Part 1 answer: 332
// Part 2 answer: 3900
func main() {
	fmt.Println("Advent of Code 2020, Day 24")
	tileList := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1. Answer = %d\n", part1(tileList))
	fmt.Printf("Part 2. Answer = %d\n", part2(tileList))
}

// Oh oh oh I remember from a previous AoC that you can use cubic coordinates for a hex grid
type coord struct {
	x, y, z int
}

var (
	e  = coord{1, -1, 0}
	se = coord{0, -1, 1}
	sw = coord{-1, 0, 1}
	w  = coord{-1, 1, 0}
	nw = coord{0, 1, -1}
	ne = coord{1, 0, -1}
)

func (c coord) add(other coord) coord {
	return coord{c.x + other.x, c.y + other.y, c.z + other.z}
}

func part1(tileList []string) int {
	return countBlackTiles(flipTiles(tileList))
}

func flipTiles(tileList []string) map[coord]bool {
	// false is white, true is black
	tiles := make(map[coord]bool)
	for _, line := range tileList {
		c := coord{}
		for _, dir := range convertLineToDirections(line) {
			c = c.add(dir)
		}
		tiles[c] = !tiles[c]
	}
	return tiles
}

func countBlackTiles(tiles map[coord]bool) int {
	var count int
	for _, v := range tiles {
		if v {
			count++
		}
	}
	return count
}

func convertLineToDirections(line string) []coord {
	var directions []coord
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case 'e':
			directions = append(directions, e)
		case 'w':
			directions = append(directions, w)
		case 'n':
			i++
			if line[i] == 'w' {
				directions = append(directions, nw)
			} else {
				directions = append(directions, ne)
			}
		case 's':
			i++
			if line[i] == 'w' {
				directions = append(directions, sw)
			} else {
				directions = append(directions, se)
			}
		}
	}
	return directions
}

func part2(tileList []string) int {
	tiles := flipTiles(tileList)
	for day := 1; day <= 100; day++ {
		tiles = applyRules(tiles)
	}
	return countBlackTiles(tiles)
}

func applyRules(tiles map[coord]bool) map[coord]bool {
	// This is a set ... bool just means membership!
	tilesToConsider := make(map[coord]bool)
	for c, black := range tiles {
		if black {
			tilesToConsider[c] = true
			for _, surrounding := range surroundingCoords(c) {
				tilesToConsider[surrounding] = true
			}
		}
	}
	// Now this is the tile map after this iteration. True means "black"
	newTiles := make(map[coord]bool)
	for c := range tilesToConsider {
		var numBlack int
		for _, surrounding := range surroundingCoords(c) {
			if tiles[surrounding] {
				numBlack++
			}
		}
		if tiles[c] {
			// Only keep it black if one or two surrounding
			if numBlack == 1 || numBlack == 2 {
				newTiles[c] = true
			}
		} else {
			// Only make it black if exactly 2 tiles
			if numBlack == 2 {
				newTiles[c] = true
			}
		}
	}
	return newTiles
}

func surroundingCoords(c coord) []coord {
	return []coord{c.add(e), c.add(se), c.add(sw), c.add(w), c.add(nw), c.add(ne)}
}
