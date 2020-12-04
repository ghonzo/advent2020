// Advent of Code 2020, Day 4
package main

import (
	"strconv"
	"strings"
	"testing"
)

func Test_passport_isValid(t *testing.T) {
	p := readPassports(strings.NewReader(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`))
	tests := []struct {
		name string
		p    passport
		want bool
	}{
		{"1", p[0], true},
		{"2", p[1], false},
		{"3", p[2], true},
		{"4", p[3], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.isValid(); got != tt.want {
				t.Errorf("passport.isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passport_isValid2_invalid(t *testing.T) {
	invalid := readPassports(strings.NewReader(`eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`))
	for i, p := range invalid {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := p.isValid2(); got {
				t.Errorf("passport.isValid2() = %v, want %v", got, false)
			}
		})
	}
}

func Test_passport_isValid2_valid(t *testing.T) {
	valid := readPassports(strings.NewReader(`pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`))
	for i, p := range valid {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := p.isValid2(); !got {
				t.Errorf("passport.isValid2() = %v, want %v", got, true)
			}
		})
	}
}
