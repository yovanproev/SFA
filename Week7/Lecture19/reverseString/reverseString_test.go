package reverseString

import (
	"io"
	"log"
	"math/rand"
	"testing"
	"time"
)

const letterBytes = "Today we test the reverse string Reader"

func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TestNewReverseStringReader(t *testing.T) {
	s := *NewReverseStringReader(RandStringBytes(10))
	// read from reverse string function
	reversedBytes, err := io.ReadAll(s)

	if err == io.EOF {
		log.Println(err)
	}

	// reverse the random string bytes
	runes := []rune(RandStringBytes(10))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	turnToStrings := string(runes)

	checkIfReverse := turnToStrings

	if string(reversedBytes) != checkIfReverse {
		t.Errorf("Incorrect, it should be %s, not %s", string(reversedBytes), checkIfReverse)
	}
}

// Ouptut
// go test ./reverseString -cover
// ok      reverseString/reverseString     0.066s  coverage: 100.0% of statements
