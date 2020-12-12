// Advent of Code 2020, Day 12
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/ghonzo/advent2020/common"
)

// Day 12: Rain Risk
// Part 1 answer: 2297
// Part 2 answer: 89984
func main() {
	fmt.Println("Advent of Code 2020, Day 12")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	instructions := readInstructions(input)
	fmt.Printf("Part 1. Answer = %d\n", part1(instructions))
	fmt.Printf("Part 2. Answer = %d\n", part2(instructions))
}

type instruction struct {
	action byte
	value  int
}

func readInstructions(r io.Reader) []instruction {
	var instructions []instruction
	input := bufio.NewScanner(r)
	for input.Scan() {
		s := input.Text()
		value, err := strconv.Atoi(s[1:])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{s[0], value})
	}
	return instructions
}

func part1(ins []instruction) int {
	var ship common.Point
	dir := common.E
	for _, i := range ins {
		switch i.action {
		case 'N':
			ship = ship.Add(common.N.Times(i.value))
		case 'S':
			ship = ship.Add(common.S.Times(i.value))
		case 'E':
			ship = ship.Add(common.E.Times(i.value))
		case 'W':
			ship = ship.Add(common.W.Times(i.value))
		case 'L':
			if i.value == 90 {
				dir = dir.Left()
			} else if i.value == 180 {
				dir = dir.Reflect()
			} else if i.value == 270 {
				dir = dir.Right()
			} else {
				panic("Bad value")
			}
		case 'R':
			if i.value == 90 {
				dir = dir.Right()
			} else if i.value == 180 {
				dir = dir.Reflect()
			} else if i.value == 270 {
				dir = dir.Left()
			} else {
				panic("Bad value")
			}
		case 'F':
			ship = ship.Add(dir.Times(i.value))
		}
	}
	return ship.ManhattanDistance()
}

func part2(ins []instruction) int {
	var ship common.Point
	waypoint := common.NewPoint(10, -1)
	for _, i := range ins {
		switch i.action {
		case 'N':
			waypoint = waypoint.Add(common.N.Times(i.value))
		case 'S':
			waypoint = waypoint.Add(common.S.Times(i.value))
		case 'E':
			waypoint = waypoint.Add(common.E.Times(i.value))
		case 'W':
			waypoint = waypoint.Add(common.W.Times(i.value))
		case 'L':
			if i.value == 90 {
				waypoint = waypoint.LeftAround(ship)
			} else if i.value == 180 {
				waypoint = waypoint.ReflectAround(ship)
			} else if i.value == 270 {
				waypoint = waypoint.RightAround(ship)
			} else {
				panic("Bad value")
			}
		case 'R':
			if i.value == 90 {
				waypoint = waypoint.RightAround(ship)
			} else if i.value == 180 {
				waypoint = waypoint.ReflectAround(ship)
			} else if i.value == 270 {
				waypoint = waypoint.LeftAround(ship)
			} else {
				panic("Bad value")
			}
		case 'F':
			move := waypoint.Sub(ship).Times(i.value)
			ship = ship.Add(move)
			waypoint = waypoint.Add(move)
		}
	}
	return ship.ManhattanDistance()
}
