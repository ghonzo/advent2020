// Advent of Code 2020, Day 18
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Day 18: Operation Order
// Part 1 answer: 29839238838303
// Part 2 answer: 201376568795521
func main() {
	fmt.Println("Advent of Code 2020, Day 18")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	lines := readInput(input)
	fmt.Printf("Part 1. Answer = %d\n", part1(lines))
	fmt.Printf("Part 2. Answer = %d\n", part2(lines))
}

func readInput(r io.Reader) []string {
	var lines []string
	input := bufio.NewScanner(r)
	for input.Scan() {
		lines = append(lines, input.Text())
	}
	return lines
}

func part1(lines []string) int {
	var sum int
	for _, s := range lines {
		sum += evalLine(s)
	}
	return sum
}

func part2(lines []string) int {
	var sum int
	for _, s := range lines {
		sum += evalLinePart2(s)
	}
	return sum
}

func evalLine(s string) int {
	pos := 0
	return eval(s, &pos)
}

func evalLinePart2(s string) int {
	return evalLine(insertParens(s))
}

// We evaluate from left to right, always increasing pos
func eval(s string, pos *int) int {
	sol := 0
	op := byte('+')
	for {
		var value int
		ch := s[*pos]
		if ch == '(' {
			*pos++
			value = eval(s, pos)
		} else {
			value = int(ch - '0')
		}
		switch op {
		case '+':
			sol += value
		case '*':
			sol *= value
		}
		*pos++
		if *pos >= len(s) || s[*pos] == ')' {
			break
		}
		*pos++
		op = s[*pos]
		*pos += 2
	}
	return sol
}

// Insert parens into the string to reflect the precedence
func insertParens(s string) string {
	for pos := 0; pos < len(s); pos++ {
		if s[pos] == '+' {
			s = parensAround(s, pos)
			pos++
		}
	}
	return s
}

// Given a string and the index of a '+', return a string that inserts parens around the two operands
func parensAround(s string, plusIndex int) string {
	var parenDepth, openParen, closeParen int
	// go backwards from the "+" to find where we should put the open paren
	for openParen = plusIndex - 2; openParen > 0; openParen-- {
		switch s[openParen] {
		case ')':
			parenDepth++
		case '(':
			parenDepth--
		}
		if parenDepth == 0 {
			break
		}
	}
	// go forwards from the "+" to find where we should put the close paren
	for closeParen = plusIndex + 2; closeParen < len(s); closeParen++ {
		switch s[closeParen] {
		case ')':
			parenDepth--
		case '(':
			parenDepth++
		}
		if parenDepth == 0 {
			break
		}
	}
	if closeParen == len(s) {
		return s[:openParen] + "(" + s[openParen:] + ")"
	}
	return s[:openParen] + "(" + s[openParen:closeParen+1] + ")" + s[closeParen+1:]
}
