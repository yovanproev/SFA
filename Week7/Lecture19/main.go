package main

import (
	"io"
	"os"
	"reverseString/reverseString"
)

func main() {
	io.Copy(os.Stdout, *reverseString.NewReverseStringReader("Hello world"))
}

// Output
// dlrow olleH
