// Advent of Code 2020, Day 7
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Day 7: Handy Haversacks
// Part 1 answer: 238
// Part 2 answer: 82930
func main() {
	fmt.Println("Advent of Code 2020, Day 7")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	rules := readRules(input)
	const target = color("shiny gold")
	fmt.Printf("Part 1. Count = %d\n", countContainedIn(target, rules))
	fmt.Printf("Part 2. Total contains = %d", countContained(target, rules)-1)
}

type color string

type colorAndQty struct {
	color
	qty int
}

var (
	containsRegex = regexp.MustCompile("(\\d+) (.*?) bag")
)

func readRules(r io.Reader) map[color][]colorAndQty {
	rules := make(map[color][]colorAndQty)
	input := bufio.NewScanner(r)
	for input.Scan() {
		ruleSides := strings.SplitN(input.Text(), " bags contain ", 2)
		key := color(ruleSides[0])
		var contains []colorAndQty
		for _, containsPart := range containsRegex.FindAllStringSubmatch((ruleSides[1]), -1) {
			// [1] is the qty, while [2] is the color
			qty, err := strconv.Atoi(containsPart[1])
			if err != nil {
				panic(err)
			}
			contains = append(contains, colorAndQty{color(containsPart[2]), qty})
		}
		rules[key] = contains
	}
	return rules
}

func colorsOnly(contents []colorAndQty) []color {
	var colors []color
	for _, caq := range contents {
		colors = append(colors, caq.color)
	}
	return colors
}

func countContainedIn(target color, rules map[color][]colorAndQty) int {
	var count int
	for bag := range rules {
		if bagDeepContains(bag, target, rules) {
			count++
		}
	}
	return count
}

func bagDeepContains(bag color, target color, rules map[color][]colorAndQty) bool {
	// This could be much more efficient if we kept intermediate results
	containsColors := colorsOnly(rules[bag])
	for len(containsColors) > 0 {
		if containsColors[0] == target {
			return true
		}
		containsColors = append(containsColors[1:], colorsOnly(rules[containsColors[0]])...)
	}
	return false
}

// Note: includes the containing bag! You may need to subtract 1 to the answer returned
func countContained(bag color, rules map[color][]colorAndQty) int {
	count := 1
	for _, caq := range rules[bag] {
		count += caq.qty * countContained(caq.color, rules)
	}
	return count
}
