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

// Collect all the rotations, strategically indexed by direction
var rotations = [4]transform{r0, r90, r180, r270}

type tile struct {
	id   int // for debugging only
	data [][]byte
	// indexed by direction
	adjacent [4]*tile
	// These all need to be initialized to "identity"
	rotation, flipV, flipH                   transform
	rotationLocked, flipVLocked, flipHLocked bool
}

var nullTile = new(tile)

func newTile() *tile {
	var t tile
	t.rotation = identity
	t.flipV = identity
	t.flipH = identity
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
	// Let's iterate over all the tiles
	for _, t := range tiles {

		t.findAdjacentTiles(tiles)
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

func (t *tile) findAdjacentTiles(tiles map[int]*tile) {
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

func (t *tile) applyAllTransforms(i, j int) (int, int) {
	return t.flipH(t.flipV(t.rotation(i, j)))
}

// This returns a bitwise representation of the edge
func (t *tile) calculateEdge(d direction) int {
	var edge int
	// Always just calculate the top edge, let the transforms do their work
	for col := 0; col < n; col++ {
		i, j := rotations[d](t.applyAllTransforms(col, 0))
		if t.data[j][i] == '#' {
			edge += (1 << col)
		}
	}
	return edge
}

func (t *tile) findEdge(edge int, d direction) bool {
	// If rotation is locked in, then let's check that
	if t.rotationLocked {
		if t.calculateEdge(d) == edge {
			// Found it, no additional transforms needed
			return true
		}
		// If we are trting to find north or south, maybe we can FlipV
		if d == north || d == south {
			if t.flipVLocked {
				// Nope, we are locked in
				return false
			}
			// Let's try
			t.flipV = flipVertical
			if t.calculateEdge(d) == edge {
				// We found it. Lock it in
				t.flipVLocked = true
				return true
			}
			// Nope, that didn't work. Set it back
			t.flipV = identity
			return false
		}
		// Since we are east or west, let's try flipH
		if t.flipHLocked {
			// Nope, we are locked in
			return false
		}
		// Let's try
		t.flipH = flipHorizontal
		if t.calculateEdge(d) == edge {
			// We found it. Lock it in
			t.flipHLocked = true
			return true
		}
		// Nope, that didn't work. Set it back
		t.flipH = identity
		return false
	}
	// We can try all rotations
	for _, rot := range rotations {
		t.rotation = rot
		if t.calculateEdge(d) == edge {
			// Got it. Lock in the correct parts
			t.rotationLocked = true
			if d == north || d == south {
				t.flipVLocked = true
			} else {
				t.flipHLocked = true
			}
			return true
		}
		// Also try a flip
		if d == north || d == south {
			t.flipV = flipVertical
			if t.calculateEdge(d) == edge {
				t.rotationLocked = true
				t.flipVLocked = true
				return true
			}
			// Nope
			t.flipV = identity
		} else {
			t.flipH = flipHorizontal
			if t.calculateEdge(d) == edge {
				t.rotationLocked = true
				t.flipHLocked = true
				return true
			}
			// Nope
			t.flipH = identity
		}
	}
	// No dice
	t.rotation = identity
	return false
}
