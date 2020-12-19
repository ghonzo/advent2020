// Advent of Code 2020, Day 19
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

// Day 19: Monster Messages
// Part 1 answer: 265
// Part 2 answer: 394
func main() {
	fmt.Println("Advent of Code 2020, Day 19")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	compiledRules, messages := readInput(input)
	fmt.Printf("Part 1. Answer = %d\n", part1(compiledRules, messages))
	fmt.Printf("Part 2. Answer = %d\n", part2(compiledRules, messages))
}

// this is the raw rule, as read from the input
type rule string

// this is a fully realized regex pattern suitable for passing to regexp.Compile()
type compiledRule string

func readInput(r io.Reader) (map[int]compiledRule, []string) {
	rules := make(map[int]rule)
	var messages []string
	input := bufio.NewScanner(r)
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break
		}
		subs := strings.Split(line, ":")
		index, err := strconv.Atoi(subs[0])
		if err != nil {
			panic(err)
		}
		rules[index] = rule(subs[1][1:])
	}
	for input.Scan() {
		messages = append(messages, input.Text())
	}
	return compile(rules), messages
}

func compile(rules map[int]rule) map[int]compiledRule {
	compiledRules := make(map[int]compiledRule)
	for ruleNum := range rules {
		compiledRules[ruleNum] = getCompiledRule(ruleNum, rules, compiledRules)
	}
	return compiledRules
}

func getCompiledRule(index int, rules map[int]rule, compiledRules map[int]compiledRule) compiledRule {
	if cr, ok := compiledRules[index]; ok {
		return cr
	}
	cr := compileRule(index, rules, compiledRules)
	compiledRules[index] = cr
	return cr
}

func compileRule(index int, rules map[int]rule, compiledRules map[int]compiledRule) compiledRule {
	// We need to compile it
	rule := rules[index]
	// literal
	if rule[0] == '"' {
		return compiledRule(rule[1:2])
	}
	var cr compiledRule
	var addEndParen bool
	for _, s := range strings.Split(string(rule), " ") {
		if s == "|" {
			// make a non-capturing group
			cr = compiledRule("(?:" + string(cr) + s)
			addEndParen = true
		} else {
			subIndex, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			cr += getCompiledRule(subIndex, rules, compiledRules)
		}
	}
	if addEndParen {
		cr += ")"
	}
	return cr
}

func part1(compiledRules map[int]compiledRule, messages []string) int {
	var count int
	for _, s := range messages {
		if compiledRules[0].matches(s) {
			count++
		}
	}
	return count
}

func (cr compiledRule) matches(message string) bool {
	// This is really inefficient that we don't cache the regexp. Oh well.
	return regexp.MustCompile("^" + string(cr) + "$").MatchString(message)
}

func part2(compiledRules map[int]compiledRule, messages []string) int {
	mm := make(map[string]bool)
	compiledRules[8] = compiledRules[42] + "+"
	// i represents the number of times we should match rule 42 and rule 31. Keep increasing until we don't match anymore.
	for i, matchedAtLeastOne := 1, true; matchedAtLeastOne; i++ {
		fmt.Println(i)
		matchedAtLeastOne = false
		repeatStr := compiledRule(fmt.Sprintf("{%d}", i))
		// Need to tweak rule0
		var cr0 compiledRule = compiledRules[8] + compiledRules[42] + repeatStr + compiledRules[31] + repeatStr
		for _, s := range messages {
			if _, ok := mm[s]; ok {
				continue
			}
			if cr0.matches(s) {
				mm[s] = true
				matchedAtLeastOne = true
			}
		}
	}
	return len(mm)
}
