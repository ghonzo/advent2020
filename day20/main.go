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

// Day 20: Jurassic Jigsaw
// Part 1 answer: 64802175715999
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2020, Day 20")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	tiles := readTiles(input)
	fmt.Printf("Part 1. Answer = %d\n", part1(tiles))
	//fmt.Printf("Part 2. Answer = %d\n", part2(compiledRules, messages))
}

type tile struct {
	data  [][]byte
	edges [8]int // each edge, also flipped
}

var tileRegexp = regexp.MustCompile(`^Tile (\d+):$`)

func readTiles(r io.Reader) map[int]tile {
	tiles := make(map[int]tile)
	input := bufio.NewScanner(r)
	for input.Scan() {
		line := input.Text()
		id, err := strconv.Atoi(tileRegexp.FindStringSubmatch(line)[1])
		if err != nil {
			panic(err)
		}
		var t tile
		for input.Scan() {
			line = input.Text()
			if line == "" {
				break
			}
			t.data = append(t.data, []byte(line))
		}
		t.calculateEdges()
		tiles[id] = t
	}
	return tiles
}

func (t *tile) calculateEdges() {
	for i, b := range t.data[0] {
		if b == '#' {
			t.edges[0] += (1 << i)
			t.edges[4] += (1 << (9 - i))
		}
	}
	for i, b := range t.data[9] {
		if b == '#' {
			t.edges[1] += (1 << i)
			t.edges[5] += (1 << (9 - i))
		}
	}
	for i := 0; i < len(t.data); i++ {
		if t.data[i][0] == '#' {
			t.edges[2] += (1 << i)
			t.edges[6] += (1 << (9 - i))
		}
	}
	for i := 0; i < len(t.data); i++ {
		if t.data[i][9] == '#' {
			t.edges[3] += (1 << i)
			t.edges[7] += (1 << (9 - i))
		}
	}
}

func part1(tiles map[int]tile) int {
	prod := 1
	for id, t := range tiles {
		matchedEdges := 0
		for _, edge := range t.edges[:4] {
			for id2, otherTile := range tiles {
				if id == id2 {
					continue
				}
				if otherTile.hasEdge(edge) {
					matchedEdges++
					break
				}
			}
		}
		fmt.Printf("Tile %d has %d matched edges %v\n", id, matchedEdges, t.edges)
		if matchedEdges == 2 {
			prod *= id
		}
	}
	return prod
}

func (t *tile) hasEdge(edge int) bool {
	for _, v := range t.edges {
		if v == edge {
			return true
		}
	}
	return false
}
