// Advent of Code 2020, Day 13
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Day 13: Shuttle Search
// Part 1 answer: 259
// Part 2 answer: 210612924879242
func main() {
	fmt.Println("Advent of Code 2020, Day 13")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	earliest, ids := readInput(input)
	fmt.Printf("Part 1. Answer = %d\n", part1(earliest, ids))
	fmt.Printf("Part 2. Answer = %d\n", part2(ids))
}

// returns earliest timestamp, and the buses. "x" is 0 in the ids
func readInput(r io.Reader) (int, []int) {
	input := bufio.NewScanner(r)
	input.Scan()
	earliest, err := strconv.Atoi(input.Text())
	if err != nil {
		panic(err)
	}
	input.Scan()
	var ids []int
	for _, idStr := range strings.Split(input.Text(), ",") {
		if idStr == "x" {
			idStr = "0"
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}
	return earliest, ids
}

func part1(earliest int, ids []int) int {
	for ts := earliest; ; ts++ {
		for _, id := range ids {
			if id != 0 && ts%id == 0 {
				return (ts - earliest) * id
			}
		}
	}
}

func part2(ids []int) uint64 {
	// We build this up id by id, incrementing the cycle time as we go
	var base uint64 = uint64(ids[0])
	inc := base
	for i, id := range ids[1:] {
		if id == 0 {
			continue
		}
		for ts := base; ; ts += inc {
			if (ts+uint64(i+1))%uint64(id) == 0 {
				// Found a new base. The new cycle is the last increment times the id we just found
				base, inc = ts, inc*uint64(id)
				break
			}
		}
	}
	return base
}
