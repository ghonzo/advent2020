package common

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// ReadInts assumes there is one integer per line and returns a slice of ints. It will panic
// if a line cannot be coverted to an int using strconv.Atoi.
func ReadInts(r io.Reader) []int {
	var ints []int
	for _, line := range ReadStrings(r) {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

// ReadIntsFromFile expects a filename that contains a list of ints, one per line. It will panic
// if there is an error opening the file.
func ReadIntsFromFile(filename string) []int {
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	return ReadInts(input)
}

// ReadStrings reads a file and returns a string slice, each strings representing a line
func ReadStrings(r io.Reader) []string {
	var strings []string
	input := bufio.NewScanner(r)
	for input.Scan() {
		strings = append(strings, input.Text())
	}
	return strings
}

// ReadStringsFromFile expects a filename and returns a slice of strings, one per file in that file.
// It will panic is there is an error opening the file.
func ReadStringsFromFile(filename string) []string {
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	return ReadStrings(input)
}
