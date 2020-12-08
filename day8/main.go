// Advent of Code 2020, Day 8
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

// Day 8: Handheld Halting
// Part 1 answer: 1594
// Part 2 answer: 758
func main() {
	fmt.Println("Advent of Code 2020, Day 8")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	instructions := readInstructions(input)
	m := machine{instructions, 0}
	runUntilDup(&m)
	fmt.Printf("Part 1. acc = %d\n", m.acc)
	m.acc = 0
	fmt.Printf("Part 2. acc = %d\n", part2(&m))
}

type instruction struct {
	operation string
	argument  int
}

var (
	instructionRegex = regexp.MustCompile("^(...) \\+?(-?\\d+)$")
)

func readInstructions(r io.Reader) []instruction {
	var instructions []instruction
	input := bufio.NewScanner(r)
	for input.Scan() {
		match := instructionRegex.FindStringSubmatch(input.Text())
		// [1] is the operation, [2] is the argument
		argument, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{match[1], argument})
	}
	return instructions
}

type machine struct {
	instructions []instruction
	acc          int
}

// Returns true if we are out of bounds, false if an instructions was executed twice
func runUntilDup(m *machine) bool {
	var ip int
	executed := make(map[int]bool)
	for {
		if ip >= len(m.instructions) || ip < 0 {
			return true
		}
		if executed[ip] {
			return false
		}
		executed[ip] = true
		switch m.instructions[ip].operation {
		case "acc":
			m.acc += m.instructions[ip].argument
			ip++
		case "jmp":
			ip += m.instructions[ip].argument
		case "nop":
			ip++
		}
	}
}

func part2(m *machine) int {
	// Starting at begining, change any "jmp" to "nop" and vice versa
	for ip := 0; ; ip++ {
		if m.instructions[ip].operation == "acc" {
			continue
		}
		// Copy the machine
		newInstructions := make([]instruction, len(m.instructions))
		copy(newInstructions, m.instructions)
		localMachine := machine{newInstructions, 0}
		switch localMachine.instructions[ip].operation {
		case "jmp":
			localMachine.instructions[ip].operation = "nop"
		case "nop":
			localMachine.instructions[ip].operation = "jmp"
		}
		if runUntilDup(&localMachine) {
			return localMachine.acc
		}
	}
}
