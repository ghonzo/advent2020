// Advent of Code 2020, Day 23
package main

import (
	"container/list"
	"fmt"
	"strconv"
)

// Day 23: Crab Cups
// Part 1 answer: 98645732
// Part 2 answer: 689500518476
func main() {
	fmt.Println("Advent of Code 2020, Day 23")
	input := "364289715"
	fmt.Printf("Part 1. Answer = %s\n", part1(input))
	fmt.Printf("Part 2. Answer = %d\n", part2(input))
}

type indexedLinkedList struct {
	cups *list.List
	// The index is cup label-1 points to where it is in the list
	cupsIndex []*list.Element
}

func newIndexedLinkedList(size int) *indexedLinkedList {
	ill := new(indexedLinkedList)
	ill.cups = list.New()
	ill.cupsIndex = make([]*list.Element, size, size)
	return ill
}

func (ill *indexedLinkedList) addCup(label int) {
	el := ill.cups.PushBack(label)
	ill.cupsIndex[label-1] = el
}

func (ill *indexedLinkedList) findCup(label int) *list.Element {
	return ill.cupsIndex[label-1]
}

func (ill *indexedLinkedList) cupAfter(cup *list.Element) *list.Element {
	after := cup.Next()
	if after == nil {
		after = ill.cups.Front()
	}
	return after
}

func (ill *indexedLinkedList) size() int {
	return len(ill.cupsIndex)
}

func (ill *indexedLinkedList) String() string {
	var s string
	for e := ill.cups.Front(); e != nil; e = e.Next() {
		s += strconv.Itoa(e.Value.(int))
	}
	return s
}

func part1(input string) string {
	ill := newIndexedLinkedList(len(input))
	for _, ch := range input {
		ill.addCup(int(ch) - '0')
	}
	currentCup := ill.cups.Front()
	for move := 0; move < 100; move++ {
		currentCup = makeMove(ill, currentCup)
	}
	// Now find all the cups after #1
	var s string
	currentCup = ill.findCup(1)
	for i := 0; i < ill.size()-1; i++ {
		currentCup = ill.cupAfter(currentCup)
		s += strconv.Itoa(currentCup.Value.(int))
	}
	return s
}

func makeMove(ill *indexedLinkedList, currentCup *list.Element) *list.Element {
	// Let's take three cups
	c1 := ill.cupAfter(currentCup)
	c2 := ill.cupAfter(c1)
	c3 := ill.cupAfter(c2)
	// Find destination cup
	destinationCupLabel := currentCup.Value.(int)
	var destinationCup *list.Element
	for {
		destinationCupLabel--
		if destinationCupLabel == 0 {
			destinationCupLabel = ill.size()
		}
		destinationCup = ill.findCup(destinationCupLabel)
		if destinationCup != c1 && destinationCup != c2 && destinationCup != c3 {
			// found it
			break
		}
	}
	// Move the removed cups there
	ill.cups.MoveAfter(c1, destinationCup)
	ill.cups.MoveAfter(c2, c1)
	ill.cups.MoveAfter(c3, c2)
	// Move clockwise
	return ill.cupAfter(currentCup)
}

func part2(input string) int {
	ill := newIndexedLinkedList(1000000)
	for _, ch := range input {
		ill.addCup(int(ch) - '0')
	}
	// Now add the rest
	for label := 10; label <= 1000000; label++ {
		ill.addCup(label)
	}
	currentCup := ill.cups.Front()
	for move := 0; move < 10000000; move++ {
		currentCup = makeMove(ill, currentCup)
	}
	// Now find the two cups after cup 1
	currentCup = ill.findCup(1)
	c1 := ill.cupAfter(currentCup)
	c2 := ill.cupAfter(c1)
	return c1.Value.(int) * c2.Value.(int)
}
