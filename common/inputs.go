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
	input := bufio.NewScanner(r)
	for input.Scan() {
		i, err := strconv.Atoi(input.Text())
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
