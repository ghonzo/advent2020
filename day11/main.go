// Advent of Code 2020, Day 11
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
)

// Day 11: Seating System
// Part 1 answer: 2310
// Part 2 answer: 2074
func main() {
	fmt.Println("Advent of Code 2020, Day 11")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	seatMap := readSeatMap(input)
	fmt.Printf("Part 1. Answer = %d\n", part1(seatMap))
	fmt.Printf("Part 2. Answer = %d\n", part2(seatMap))
}

type seatMap [][]byte

func readSeatMap(r io.Reader) seatMap {
	var sm seatMap
	input := bufio.NewScanner(r)
	for input.Scan() {
		sm = append(sm, []byte(input.Text()))
	}
	return sm
}

func (sm seatMap) sizeY() int {
	return len(sm)
}

func (sm seatMap) sizeX() int {
	return len(sm[0])
}

// if second argument (ok) = true, then in bounds
func (sm seatMap) get(x, y int) (byte, bool) {
	if x < 0 || x >= sm.sizeX() || y < 0 || y >= sm.sizeY() {
		return floor, false
	}
	return sm[y][x], true
}

func (sm seatMap) set(x, y int, b byte) {
	sm[y][x] = b
}

func (sm seatMap) countOccupied() int {
	var count int
	for y := 0; y < sm.sizeY(); y++ {
		for x := 0; x < sm.sizeX(); x++ {
			if v, _ := sm.get(x, y); v == occupied {
				count++
			}
		}
	}
	return count
}

func part1(sm seatMap) int {
	for i := 0; ; i++ {
		//fmt.Println("Cycle ", i, " occupied ", sm.countOccupied())
		newSm := sm.cycle()
		// cheap shot
		if reflect.DeepEqual(sm, newSm) {
			return sm.countOccupied()
		}
		sm = newSm
	}
}

func part2(sm seatMap) int {
	for i := 0; ; i++ {
		//fmt.Println("Cycle ", i, " occupied ", sm.countOccupied())
		newSm := sm.cycle2()
		if reflect.DeepEqual(sm, newSm) {
			return sm.countOccupied()
		}
		sm = newSm
	}
}

const (
	occupied = '#'
	empty    = 'L'
	floor    = '.'
)

func (sm seatMap) cycle() seatMap {
	retVal := make(seatMap, sm.sizeY())
	for y := 0; y < sm.sizeY(); y++ {
		retVal[y] = make([]byte, sm.sizeX())
		for x := 0; x < sm.sizeX(); x++ {
			retVal.set(x, y, sm.applyRule(x, y))
		}
	}
	return retVal
}

func (sm seatMap) cycle2() seatMap {
	retVal := make(seatMap, sm.sizeY())
	for y := 0; y < sm.sizeY(); y++ {
		retVal[y] = make([]byte, sm.sizeX())
		for x := 0; x < sm.sizeX(); x++ {
			retVal.set(x, y, sm.applyRule2(x, y))
		}
	}
	return retVal
}

func (sm seatMap) occupiedAround(x, y int) int {
	var o int
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if v, _ := sm.get(i, j); v == occupied {
				o++
			}
		}
	}
	return o
}

func (sm seatMap) occupiedAround2(x, y int) int {
	var o int
	o += sm.look(x, y, -1, -1)
	o += sm.look(x, y, 0, -1)
	o += sm.look(x, y, 1, -1)
	o += sm.look(x, y, -1, 0)
	o += sm.look(x, y, 1, 0)
	o += sm.look(x, y, -1, 1)
	o += sm.look(x, y, 0, 1)
	o += sm.look(x, y, 1, 1)
	return o
}

// return 1 if occupied, 0 else
func (sm seatMap) look(x, y, dx, dy int) int {
	for {
		x += dx
		y += dy
		if v, ok := sm.get(x, y); !ok {
			// walked off the edge
			return 0
		} else if v == occupied {
			return 1
		} else if v == empty {
			return 0
		}
		// must be floor ... keep going
	}
}

func (sm seatMap) applyRule(x, y int) byte {
	current, _ := sm.get(x, y)
	o := sm.occupiedAround(x, y)
	if current == empty && o == 0 {
		return occupied
	}
	if current == occupied && o >= 4 {
		return empty
	}
	return current
}

func (sm seatMap) applyRule2(x, y int) byte {
	current, _ := sm.get(x, y)
	o := sm.occupiedAround2(x, y)
	if current == empty && o == 0 {
		return occupied
	}
	if current == occupied && o >= 5 {
		return empty
	}
	return current
}
