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
	rulesMap := make(map[color]rule)
	for _, rule := range rules {
		rulesMap[rule.color] = rule
	}
	var count int
	for _, rule := range rules {
		if bagDeepContains(rule.color, color("shiny gold"), rulesMap) {
			count++
		}
	}
	fmt.Printf("Part 1. Count = %d\n", count)
	fmt.Printf("Part 2. Total contains = %d", countContained(color("shiny gold"), rulesMap)-1)
}

func bagDeepContains(ruleColor color, targetColor color, rulesMap map[color]rule) bool {
	containsColors := rulesMap[ruleColor].containsColors()
	for len(containsColors) > 0 {
		if containsColors[0] == targetColor {
			return true
		}
		containsColors = append(containsColors[1:], rulesMap[containsColors[0]].containsColors()...)
	}
	return false
}

func countContained(bag color, rulesMap map[color]rule) int {
	count := 1
	for _, caq := range rulesMap[bag].contains {
		count += caq.qty * countContained(caq.color, rulesMap)
	}
	return count
}

type color string

type colorAndQty struct {
	color
	qty int
}

type rule struct {
	color
	contains []colorAndQty
}

func (r rule) containsColors() []color {
	var colors []color
	for _, caq := range r.contains {
		colors = append(colors, caq.color)
	}
	return colors
}

var (
	containsRegex = regexp.MustCompile("(\\d+) (.*?) bag")
)

func readRules(r io.Reader) []rule {
	var rules []rule
	input := bufio.NewScanner(r)
	for input.Scan() {
		line := input.Text()
		var r rule
		// could use Split
		bagsIndex := strings.Index(line, " bags")
		r.color = color(line[:bagsIndex])
		rightSide := line[bagsIndex:]
		for _, containsPart := range containsRegex.FindAllStringSubmatch((rightSide), -1) {
			// [1] is the qty, while [2] is the color
			qty, err := strconv.Atoi(containsPart[1])
			if err != nil {
				panic(err)
			}
			c := containsPart[2]
			r.contains = append(r.contains, colorAndQty{color(c), qty})
		}
		rules = append(rules, r)
	}
	return rules
}

func (r rule) String() string {
	return fmt.Sprintf("%s bags contain %v", r.color, r.contains)
}
