// Advent of Code 2020, Day 16
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

// Day 16: Ticket Translation
// Part 1 answer: 23009
// Part 2 answer: 10458887314153
func main() {
	fmt.Println("Advent of Code 2020, Day 16")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	rules, yourTicket, nearbyTickets := readInput(input)
	fmt.Printf("Part 1. Answer = %d\n", part1(rules, nearbyTickets))
	fmt.Printf("Part 2. Answer = %d\n", part2(rules, yourTicket, nearbyTickets))
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

var ruleRegexp = regexp.MustCompile(`^([^:]+): (\d+)-(\d+) or (\d+)-(\d+)$`)

type rule struct {
	name                   string
	min1, max1, min2, max2 int
}

func (r rule) isValid(field int) bool {
	return (field >= r.min1 && field <= r.max1) || (field >= r.min2 && field <= r.max2)
}

type ticket []int

func readInput(r io.Reader) (rules []rule, yourTicket ticket, nearbyTickets []ticket) {
	input := bufio.NewScanner(r)
	// Rules
	for input.Scan() {
		s := input.Text()
		// blank line means go to next section
		if s == "" {
			break
		}
		m := ruleRegexp.FindStringSubmatch(s)
		rules = append(rules, rule{m[1], atoi(m[2]), atoi(m[3]), atoi(m[4]), atoi(m[5])})
	}
	input.Scan() // "your ticket:"
	input.Scan()
	yourTicket = parseTicket(input.Text())
	input.Scan() // blank
	input.Scan() // "nearby tickets:"
	for input.Scan() {
		nearbyTickets = append(nearbyTickets, parseTicket(input.Text()))
	}
	// bare return
	return
}

func parseTicket(s string) ticket {
	var t ticket
	for _, fieldStr := range strings.Split(s, ",") {
		t = append(t, atoi(fieldStr))
	}
	return t
}

func part1(rules []rule, nearbyTickets []ticket) int {
	var sum int
	for _, t := range nearbyTickets {
		for _, field := range t {
			if !validField(rules, field) {
				sum += field
			}
		}
	}
	return sum
}

// validField returns true if the field is valid according to at least one rule
func validField(rules []rule, field int) bool {
	for _, r := range rules {
		if r.isValid(field) {
			return true
		}
	}
	return false
}

// validTicket returns true if all the fields are valid
func validTicket(rules []rule, t ticket) bool {
	for _, field := range t {
		if !validField(rules, field) {
			return false
		}
	}
	return true
}

func part2(rules []rule, yourTicket ticket, nearbyTickets []ticket) int {
	rulesInFieldOrder := findRulesInFieldOrder(rules, append(nearbyTickets, yourTicket))
	// Now we are left with just the rules
	answer := 1
	for fieldIndex, r := range rulesInFieldOrder {
		if strings.HasPrefix(r.name, "departure") {
			answer *= yourTicket[fieldIndex]
		}
	}
	return answer
}

func findRulesInFieldOrder(rules []rule, tickets []ticket) []rule {
	// We start with every field holding onto all possible rules, then eliminate them.
	// The index of ruleMaps matches the field index
	var ruleMaps []map[rule]bool
	for range rules {
		ruleMaps = append(ruleMaps, rulesAsMap(rules))
	}
	// Now for each field on each ticket, eliminate rules that don't belong
	for _, t := range tickets {
		if !validTicket(rules, t) {
			// Throw out invalid tickets
			continue
		}
		for fieldIndex, field := range t {
			for _, r := range rules {
				if !r.isValid(field) {
					// This rule is not valid for this field, so eliminate the rule
					delete(ruleMaps[fieldIndex], r)
				}
			}
		}
	}
	// We have the set of rules that are valid for each field. Find the field that has just one valid rule
	// and eliminate it from all other fields, and repeat until each field has just one valid rule
	return reduceRuleMaps(ruleMaps)
}

// Warning: this obliterates the map that's passed in
func reduceRuleMaps(ruleMaps []map[rule]bool) []rule {
	rulesInFieldOrder := make([]rule, len(ruleMaps))
	found := 0
	for found < len(ruleMaps) {
		for fieldIndex, ruleMap := range ruleMaps {
			if rulesInFieldOrder[fieldIndex].name != "" {
				// Already found, so skip
				continue
			}
			if len(ruleMap) == 1 {
				// Found the rule ... remove it from all the other possibilities
				r := getOnlyRule(ruleMap)
				rulesInFieldOrder[fieldIndex] = r
				found++
				for _, rm := range ruleMaps {
					delete(rm, r)
				}
			}
		}
	}
	return rulesInFieldOrder
}

func rulesAsMap(rules []rule) map[rule]bool {
	ruleMap := make(map[rule]bool)
	for _, r := range rules {
		ruleMap[r] = true
	}
	return ruleMap
}

func getOnlyRule(ruleMap map[rule]bool) rule {
	if len(ruleMap) > 1 {
		panic("More than one rule")
	}
	for key := range ruleMap {
		return key
	}
	// never get here
	panic("Should never get here")
}
