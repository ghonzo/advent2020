// Advent of Code 2020, Day 22
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Day 22: Crab Combat
// Part 1 answer: 30197
// Part 2 answer: 34031
func main() {
	fmt.Println("Advent of Code 2020, Day 22")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	decks := readDecks(input)
	fmt.Printf("Part 1. Answer = %d\n", part1(decks))
	fmt.Printf("Part 2. Answer = %d\n", part2(decks))
}

type deck []byte

func readDecks(r io.Reader) [2]deck {
	var decks [2]deck
	currentDeck := 0
	input := bufio.NewScanner(r)
	input.Scan()
	for input.Scan() {
		line := input.Text()
		if line == "" {
			currentDeck++
			input.Scan()
			continue
		}
		v, err := strconv.Atoi(input.Text())
		if err != nil {
			panic(err)
		}
		decks[currentDeck] = append(decks[currentDeck], byte(v))
	}
	return decks
}

func part1(arg [2]deck) int {
	var decks [2]deck
	decks[0] = make(deck, len(arg[0]))
	copy(decks[0], arg[0])
	decks[1] = make(deck, len(arg[1]))
	copy(decks[1], arg[1])
	for len(decks[0]) > 0 && len(decks[1]) > 0 {
		card0 := decks[0][0]
		decks[0] = decks[0][1:]
		card1 := decks[1][0]
		decks[1] = decks[1][1:]
		if card0 > card1 {
			decks[0] = append(decks[0], card0, card1)
		} else {
			decks[1] = append(decks[1], card1, card0)
		}
	}
	return score(decks)
}

func score(decks [2]deck) int {
	d := decks[0]
	if len(d) == 0 {
		d = decks[1]
	}
	var sum int
	for i, card := range d {
		sum += int(card) * (len(d) - i)
	}
	return sum
}

func part2(arg [2]deck) int {
	var decks [2]deck
	decks[0] = make(deck, len(arg[0]))
	copy(decks[0], arg[0])
	decks[1] = make(deck, len(arg[1]))
	copy(decks[1], arg[1])
	_, finalDecks := playGame(decks, [2]byte{byte(len(decks[0])), byte(len(decks[1]))})
	return score(finalDecks)
}

func playGame(decks [2]deck, cards [2]byte) (int, [2]deck) {
	previousStates := make(map[string]bool)
	var subDecks [2]deck
	subDecks[0] = make(deck, cards[0])
	copy(subDecks[0], decks[0][:cards[0]])
	subDecks[1] = make(deck, cards[1])
	copy(subDecks[1], decks[1][:cards[1]])
	var winningSubDeck int
	for len(subDecks[0]) > 0 && len(subDecks[1]) > 0 {
		if str := decksState(subDecks); previousStates[str] {
			// Winner by recursion
			return 0, subDecks
		} else {
			previousStates[str] = true
		}
		var c [2]byte
		c[0] = subDecks[0][0]
		subDecks[0] = subDecks[0][1:]
		c[1] = subDecks[1][0]
		subDecks[1] = subDecks[1][1:]
		if len(subDecks[0]) >= int(c[0]) && len(subDecks[1]) >= int(c[1]) {
			winningSubDeck, _ = playGame(subDecks, c)
		} else if c[0] > c[1] {
			winningSubDeck = 0
		} else {
			winningSubDeck = 1
		}
		subDecks[winningSubDeck] = append(subDecks[winningSubDeck], c[winningSubDeck], c[1-winningSubDeck])
	}
	return winningSubDeck, subDecks
}

func decksState(decks [2]deck) string {
	return string(decks[0]) + "|" + string(decks[1])
}
