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
// Part 1 answer:
// Part 2 answer:
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
	//fmt.Printf("Rule 8 = %s\n", compiledRules[8])
	//fmt.Printf("Rule 11 = %s\n", compiledRules[11])
	//fmt.Printf("Rule 42 = %s\n", compiledRules[42])
	//fmt.Printf("Rule 31 = %s\n", compiledRules[31])
}

type rule string

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
	compiledRules := make(map[int]compiledRule)
	for k := range rules {
		compiledRules[k] = getCompiledRule(k, rules, compiledRules)
	}
	for input.Scan() {
		messages = append(messages, input.Text())
	}
	return compiledRules, messages
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
			cr = compiledRule("(?:" + string(cr) + s)
			addEndParen = true
		} else {
			subIndex, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			cr = cr + getCompiledRule(subIndex, rules, compiledRules)
		}
	}
	if addEndParen {
		cr = cr + ")"
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
	return regexp.MustCompile("^" + string(cr) + "$").MatchString(message)
}

func part2(compiledRules map[int]compiledRule, messages []string) int {
	mm := make(map[string]bool)
	compiledRules[8] = compiledRules[42] + "+"
	for i, matchedAtLeastOne := 1, true; matchedAtLeastOne; i++ {
		fmt.Println(i)
		matchedAtLeastOne = false
		repeatStr := fmt.Sprintf("{%d}", i)
		// Need to tweak rule0
		cr0 := compiledRule(string(compiledRules[8]) + string(compiledRules[42]) + repeatStr + string(compiledRules[31]) + repeatStr)
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
