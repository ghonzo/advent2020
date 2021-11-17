// Advent of Code 2020, Day 2
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2020/common"
)

// Day 2: Password Philosophy
// Part 1 answer: 474
// Part 2 answer: 745
func main() {
	fmt.Println("Advent of Code 2020, Day 2")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	policies := readPasswordPolicies(input)
	var numValid int
	for _, p := range policies {
		if p.valid() {
			numValid++
		}
	}
	fmt.Printf("Part 1: valid entries =  %d\n", numValid)
	numValid = 0
	for _, p := range policies {
		if p.valid2() {
			numValid++
		}
	}
	fmt.Printf("Part 2: valid entries =  %d\n", numValid)
}

type passwordPolicy struct {
	Min, Max int
	Letter   byte
	Password string
}

func readPasswordPolicies(r io.Reader) []passwordPolicy {
	var policies []passwordPolicy
	for _, line := range common.ReadStrings(r) {
		var pp passwordPolicy
		var err error
		hyphenIndex := strings.Index(line, "-")
		spaceIndex := strings.Index(line, " ")
		pp.Min, err = strconv.Atoi(line[:hyphenIndex])
		if err != nil {
			fmt.Printf("%s: %s\n", err, line)
		}
		pp.Max, err = strconv.Atoi(line[hyphenIndex+1 : spaceIndex])
		if err != nil {
			fmt.Printf("%s: %s\n", err, line)
		}
		pp.Letter = line[spaceIndex+1]
		pp.Password = line[spaceIndex+4:]
		policies = append(policies, pp)
	}
	return policies
}

func (p passwordPolicy) valid() bool {
	numLetters := strings.Count(p.Password, string(p.Letter))
	return numLetters >= p.Min && numLetters <= p.Max
}

func (p passwordPolicy) valid2() bool {
	var matches int
	if p.Password[p.Min-1] == p.Letter {
		matches++
	}
	if p.Password[p.Max-1] == p.Letter {
		matches++
	}
	return matches == 1
}
