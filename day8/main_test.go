// Advent of Code 2020, Day 8
package main

import (
	"strings"
	"testing"
)

func Test_runUntilDup(t *testing.T) {
	instructions := readInstructions(strings.NewReader(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`))
	m := machine{instructions, 0}
	runUntilDup(&m)
	if m.acc != 5 {
		t.Errorf("runUntilDup() = %v, want 5", m.acc)
	}
}

func Test_part2(t *testing.T) {
	instructions := readInstructions(strings.NewReader(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`))
	m := machine{instructions, 0}
	acc := part2(&m)
	if acc != 8 {
		t.Errorf("part2() = %v, want 8", acc)
	}
}
