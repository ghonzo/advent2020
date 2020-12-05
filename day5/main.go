// Advent of Code 2020, Day 5
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Day 5: Binary Boarding
// Part 1 answer: 864
// Part 2 answer: 739
func main() {
	fmt.Println("Advent of Code 2020, Day 5")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	boardingPasses := readBoardingPasses(input)
	var highest int
	takenIds := make(map[int]bool)
	for _, b := range boardingPasses {
		s := b.seatID()
		takenIds[s] = true
		if s > highest {
			highest = s
		}
	}
	fmt.Printf("Part 1: highest =  %d\n", highest)
	for i := 9; ; i++ {
		if !takenIds[i] && takenIds[i-1] && takenIds[i+1] {
			fmt.Printf("Part 2: seatId =  %d\n", i)
			break
		}
	}
}

type boardingPass string

func readBoardingPasses(r io.Reader) []boardingPass {
	var boardingPasses []boardingPass
	input := bufio.NewScanner(r)
	for input.Scan() {
		boardingPasses = append(boardingPasses, boardingPass(input.Text()))
	}
	return boardingPasses
}

func (b boardingPass) row() int {
	i, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(string(b)[:7], "F", "0"), "B", "1"), 2, 0)
	return int(i)
}

func (b boardingPass) column() int {
	i, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(string(b)[7:], "L", "0"), "R", "1"), 2, 0)
	return int(i)
}

func (b boardingPass) seatID() int {
	return b.row()*8 + b.column()
}
