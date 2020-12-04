// Advent of Code 2020, Day 4
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

// Day 4: Passport Processing
// Part 1 answer: 239
// Part 2 answer: 188
func main() {
	fmt.Println("Advent of Code 2020, Day 4")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	passports := readPassports(input)
	var valid, valid2 int
	for _, p := range passports {
		if p.isValid() {
			valid++
		}
		if p.isValid2() {
			valid2++
		}
	}
	fmt.Printf("Part 1: valid =  %d\n", valid)
	fmt.Printf("Part 2: valid =  %d\n", valid2)
}

type passport map[string]string

func readPassports(r io.Reader) []passport {
	var passports []passport
	input := bufio.NewScanner(r)
	p := make(passport)
	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			passports = append(passports, p)
			p = make(passport)
		} else {
			for _, e := range strings.Split(line, " ") {
				colonIndex := strings.Index(e, ":")
				p[e[:colonIndex]] = e[colonIndex+1:]
			}
		}
	}
	return append(passports, p)
}

func (p passport) isValid() bool {
	var requiredFields = [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range requiredFields {
		if _, ok := p[field]; !ok {
			return false
		}
	}
	return true
}

func (p passport) isValid2() bool {
	hclPattern := regexp.MustCompile("^#[0-9a-f]{6}$")
	eclPattern := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	pidPattern := regexp.MustCompile("^\\d{9}$")

	if v, ok := p["byr"]; !ok {
		return false
	} else if byr, err := strconv.Atoi(v); err != nil || byr < 1920 || byr > 2002 {
		return false
	}
	if v, ok := p["iyr"]; !ok {
		return false
	} else if iyr, err := strconv.Atoi(v); err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}
	if v, ok := p["eyr"]; !ok {
		return false
	} else if eyr, err := strconv.Atoi(v); err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}
	if v, ok := p["hgt"]; !ok {
		return false
	} else if strings.HasSuffix(v, "cm") && len(v) == 5 {
		if hgt, err := strconv.Atoi(v[:3]); err != nil || hgt < 150 || hgt > 193 {
			return false
		}
	} else if strings.HasSuffix(v, "in") && len(v) == 4 {
		if hgt, err := strconv.Atoi(v[:2]); err != nil || hgt < 59 || hgt > 76 {
			return false
		}
	} else {
		// not in cm or in
		return false
	}
	if v, ok := p["hcl"]; !ok || !hclPattern.MatchString(v) {
		return false
	}
	if v, ok := p["ecl"]; !ok || !eclPattern.MatchString(v) {
		return false
	}
	if v, ok := p["pid"]; !ok || !pidPattern.MatchString(v) {
		return false
	}
	return true
}
