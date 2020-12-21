// Advent of Code 2020, Day 20
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

// Day 20: Jurassic Jigsaw (Take 2)
// Part 1 answer: 64802175715999
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2020, Day 20b")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	tiles := readTiles(input)
	fmt.Printf("Part 1. Answer = %d\n", part1(tiles))
}

type transform func(i, j int) (int, int)

// size of the tile
const n = 10

// TRANSFORMS

func identity(i, j int) (int, int) {
	return i, j
}

// Rotations
func r0(i, j int) (int, int) {
	return identity(i, j)
}

func r90(i, j int) (int, int) {
	return n - j - 1, i
}

func r180(i, j int) (int, int) {
	return n - i - 1, n - j - 1
}

func r270(i, j int) (int, int) {
	return j, n - i - 1
}

// Flips
func flipVertical(i, j int) (int, int) {
	return n - i - 1, j
}

func flipHorizontal(i, j int) (int, int) {
	return i, n - j - 1
}

// Composites
func r90FlipVertical(i, j int) (int, int) {
	return flipVertical(r90(i, j))
}

func r90FlipHorizontal(i, j int) (int, int) {
	return flipHorizontal(r90(i, j))
}

type direction int

const (
	north direction = iota
	east
	south
	west
)

var directions = []direction{north, east, south, west}

func (d direction) opposite() direction {
	return direction((d + 2) % 4)
}

// Collect all the transforms. The first four are significant because indexed by direction
var allTransforms = [8]transform{r0, r90, r180, r270, flipVertical, flipHorizontal, r90FlipVertical, r90FlipHorizontal}

type tile struct {
	id   int // for debugging only
	data [][]byte
	// indexed by direction
	adjacent [4]*tile
	xform    transform
	locked   bool
}

var nullTile = new(tile)

func newTile() *tile {
	var t tile
	t.xform = identity
	return &t
}

var tileRegexp = regexp.MustCompile(`^Tile (\d+):$`)

func readTiles(r io.Reader) map[int]*tile {
	tiles := make(map[int]*tile)
	input := bufio.NewScanner(r)
	for input.Scan() {
		line := input.Text()
		id, err := strconv.Atoi(tileRegexp.FindStringSubmatch(line)[1])
		if err != nil {
			panic(err)
		}
		t := newTile()
		t.id = id
		for input.Scan() {
			line = input.Text()
			if line == "" {
				break
			}
			t.data = append(t.data, []byte(line))
		}
		tiles[id] = t
	}
	return tiles
}

// This will not only return the product of the corners, but will
// leave all the tiles appropriately rotated, flipped, and connected
func part1(tiles map[int]*tile) int {
	//var tilesToSearch = []*tile{tiles[3079]}
	var tilesToSearch []*tile
	// Pick a tile, any tile
	for _, t := range tiles {
		tilesToSearch = []*tile{t}
	}

	// Let's iterate over all the tiles to search until they are all found
	for len(tilesToSearch) > 0 {
		t := tilesToSearch[0]
		tilesToSearch = tilesToSearch[1:]
		if !t.allFound() {
			tilesToSearch = append(tilesToSearch, t.findAdjacentTiles(tiles)...)
		}
	}
	// Okay, now that they are all placed, find the corners
	prod := 1
	for id, t := range tiles {
		fmt.Printf("Tile id %d has %d adjacent\n", id, t.countAdjacent())
		if t.countAdjacent() == 2 {
			prod *= id
		}
	}
	return prod
}

func (t *tile) allFound() bool {
	if t == nullTile {
		return true
	}
	for _, v := range t.adjacent {
		if v == nil {
			return false
		}
	}
	return true
}

func (t *tile) findAdjacentTiles(tiles map[int]*tile) []*tile {
	t.locked = true
DirectionSearch:
	for _, d := range directions {
		if t.adjacent[d] != nil {
			// Already have that direction
			continue
		}
		edge := t.calculateEdge(d)
		// Now look at all other tiles to see if the edge is there
		for _, otherTile := range tiles {
			// But not me!
			if otherTile == t {
				continue
			}
			if otherTile.findEdge(edge, d.opposite()) {
				t.adjacent[d] = otherTile
				otherTile.adjacent[d.opposite()] = t
				continue DirectionSearch
			}
		}
		// Not found, so put nullTile in there
		t.adjacent[d] = nullTile
	}
	return t.adjacent[:]
}

func (t *tile) countAdjacent() int {
	var count int
	for _, v := range t.adjacent {
		if v != nullTile {
			count++
		}
	}
	return count
}

// This returns a bitwise representation of the edge
func (t *tile) calculateEdge(d direction) int {
	var edge int
	// Always just calculate the top edge, let the transforms do their work
	for col := 0; col < n; col++ {
		i, j := t.xform(allTransforms[d](col, 0))
		if t.data[j][i] == '#' {
			edge += (1 << col)
		}
	}
	return edge
}

// This returns a bitwise representation of the edge for matching
func (t *tile) calculateEdgeReversed(d direction) int {
	var edge int
	// Always just calculate the top edge, let the transforms do their work
	for col := 0; col < n; col++ {
		i, j := t.xform(allTransforms[d](n-col-1, 0))
		if t.data[j][i] == '#' {
			edge += (1 << col)
		}
	}
	return edge
}
func (t *tile) findEdge(edge int, d direction) bool {
	// If rotation is locked in, then let's check that
	if t.locked {
		return t.calculateEdgeReversed(d) == edge
	}
	// We can try all transforms
	for _, xf := range allTransforms {
		t.xform = xf
		if t.calculateEdgeReversed(d) == edge {
			// Got it. Lock it in
			t.locked = true
			return true
		}
	}
	// No dice (not sure if we need to set this back but whatevs)
	t.xform = identity
	return false
}
