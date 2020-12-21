// Advent of Code 2020, Day 20
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

// Day 20: Jurassic Jigsaw (Take 2)
// Part 1 answer: 64802175715999
// Part 2 answer: 2146
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
	fmt.Printf("Part 2. Answer = %d\n", part2(tiles))
}

type transform func(i, j int) (int, int)

// size of the tile
var n = 10

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

func (t *tile) get(i, j int) byte {
	ti, tj := t.xform(i, j)
	return t.data[tj][ti]
}

var monster = [3][]byte{
	[]byte("                  # "),
	[]byte("#    ##    ##    ###"),
	[]byte(" #  #  #  #  #  #   "),
}

const monsterCovers = 15

func part2(tiles map[int]*tile) int {
	// There is probably a fancier way to do this, but I'm going to copy all the subtiles into one big tile
	var bigTileData [][]byte
	rowStartTile := findNWTile(tiles)
	var bigTileDataRow int
	for t := rowStartTile; t != nullTile; {
		for j := 1; j < n-1; j++ {
			var row []byte
			for i := 1; i < n-1; i++ {
				row = append(row, t.get(i, j))
			}
			bigTileData = append(bigTileData, row)
		}
		for t = t.adjacent[east]; t != nullTile; t = t.adjacent[east] {
			for j := 1; j < n-1; j++ {
				for i := 1; i < n-1; i++ {
					bigTileData[bigTileDataRow+j-1] = append(bigTileData[bigTileDataRow+j-1], t.get(i, j))
				}
			}
		}
		// Advance to the next "tile row"
		rowStartTile = rowStartTile.adjacent[south]
		t = rowStartTile
		bigTileDataRow += (n - 2)
	}
	// Okay now we have a big huge tile
	for _, row := range bigTileData {
		fmt.Println(string(row))
	}
	// Count all the #
	var count int
	for _, row := range bigTileData {
		count += bytes.Count(row, []byte("#"))
	}
	// Let's define a big tile for our big tile data
	bigTile := newTile()
	bigTile.data = bigTileData
	// Oh dear God please forgive this hack
	n = len(bigTileData)
	// Okay let's go hunting for monsters
	var monsters int
	for _, xf := range allTransforms {
		bigTile.xform = xf
		// Now define a moving window to find the monster
		for j := 0; j <= n-len(monster); j++ {
			for i := 0; i <= n-len(monster[0]); i++ {
				if detectMonster(bigTile, i, j) {
					monsters++
				}
			}
		}
		if monsters > 0 {
			fmt.Printf("Found %d monsters for xform %v\n", monsters, xf)
			return count - monsters*15 // 15 is the number of lit pixels in the monster
		}
	}
	panic("No monsters found")
}

func findNWTile(tiles map[int]*tile) *tile {
	for _, t := range tiles {
		if t.adjacent[north] == nullTile && t.adjacent[west] == nullTile {
			return t
		}
	}
	panic("No NW tile found")
}

func detectMonster(t *tile, i, j int) bool {
	for mj, mrow := range monster {
		for mi, pixel := range mrow {
			if pixel == '#' && t.get(i+mi, j+mj) != '#' {
				return false
			}
		}
	}
	return true
}
