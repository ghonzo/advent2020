// Advent of Code 2020, Day 14
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2020/common"
)

// Day 14: Docking Data
// Part 1 answer: 7611244640053
// Part 2 answer: 3705162613854
func main() {
	fmt.Println("Advent of Code 2020, Day 14")
	instructions := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1. Answer = %d\n", part1(instructions))
	fmt.Printf("Part 2. Answer = %d\n", part2(instructions))
}

var mem = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func part1(instructions []string) uint64 {
	address := make(map[int]uint64)
	var mask string
	for _, s := range instructions {
		if strings.HasPrefix(s, "mask") {
			mask = s[7:]
		} else {
			addr, v := parseMem(s)
			address[addr] = applyMask(v, mask)
		}
	}
	var sum uint64
	for _, v := range address {
		sum += v
	}
	return sum
}

func parseMem(instruction string) (addr int, v uint64) {
	parts := mem.FindStringSubmatch(instruction)
	addr, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	v, err = strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		panic(err)
	}
	return
}

func parseBinary(s string) uint64 {
	b, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return b
}

func applyMask(v uint64, mask string) uint64 {
	// First find the value to OR with. That's all the 1s
	orMask := parseBinary(strings.ReplaceAll(mask, "X", "0"))
	// Now find the value to do AND with. That's all the 0s
	andMask := parseBinary(strings.ReplaceAll(mask, "X", "1"))
	return (v | orMask) & andMask
}

func part2(instructions []string) uint64 {
	address := make(map[uint64]uint64)
	var mask string
	for _, s := range instructions {
		if strings.HasPrefix(s, "mask") {
			mask = s[7:]
		} else {
			addr, v := parseMem(s)
			for _, a := range getAllAddresses(uint64(addr), mask) {
				address[a] = v
			}
		}
	}
	var sum uint64
	for _, v := range address {
		sum += v
	}
	return sum
}

func getAllAddresses(addr uint64, maskStr string) []uint64 {
	var masks, addresses []uint64
	addr = adjustBaseAddress(addr, maskStr)
	findAllMasks(maskStr, &masks)
	for _, mask := range masks {
		addresses = append(addresses, addr|mask)
	}
	return addresses
}

func adjustBaseAddress(addr uint64, maskStr string) uint64 {
	//First I have to adjust the address ... OR it with X replaced with 0
	orMask := parseBinary(strings.ReplaceAll(maskStr, "X", "0"))
	// Now do and AND NOT, replace with 1 -> 0 and X -> 1
	str := strings.Map(func(r rune) rune {
		switch r {
		case '1':
			return '0'
		case 'X':
			return '1'
		default:
			return r
		}
	}, maskStr)
	andNotMask := parseBinary(str)
	return (addr | orMask) &^ andNotMask
}

func findAllMasks(maskStr string, masks *[]uint64) {
	if strings.IndexByte(maskStr, 'X') < 0 {
		*masks = append(*masks, parseBinary(maskStr))
	} else {
		findAllMasks(strings.Replace(maskStr, "X", "0", 1), masks)
		findAllMasks(strings.Replace(maskStr, "X", "1", 1), masks)
	}
}
