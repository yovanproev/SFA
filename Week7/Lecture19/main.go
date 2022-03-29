package main

import (
	"io"
	"os"
	"strings"
)

type ReverseStringReader io.Reader

func main() {
	io.Copy(os.Stdout, *NewReverseStringReader("Hello world"))
}

func NewReverseStringReader(input string) *ReverseStringReader {
	var reader ReverseStringReader

	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	turnToStrings := string(runes)

	toReader := strings.NewReader(turnToStrings)

	reader = toReader
	return &reader
}

// Output
// dlrow olleH
