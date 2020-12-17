// Advent of Code 2020, Day 17
package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_cycle(t *testing.T) {
	s := readState(strings.NewReader(".#.\n..#\n###"))
	fmt.Printf("Initial: %d\n", len(s))
	for i := 0; i < 6; i++ {
		s := cycle(s)
		fmt.Printf("Cycle %d: %d\n", i, len(s))
	}
}
