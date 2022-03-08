package main

import (
	"errors"
	"fmt"
	"log"
)

type Action func() error

func main() {
	//err := SafeExec(funcWithError, "Panic")
	err := SafeExec(funcWithError, "")

	if err != nil {
		log.Fatalf("there was an error: %v", err())
	}
}

var (
	funcWithError = func() error {
		return errors.New("Error producing function")
	}
)

func SafeExec(a Action, state string) Action {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("safe exec: ", r)
		}
	}()

	if state == "Panic" {
		panic("state cannot be Panic, panicking!")
	}
	return funcWithError
}

// Input: err := SafeExec(funcWithError, "Panic")
// Output: safe exec:  state cannot be Panic, panicking!

// Input: err := SafeExec(funcWithError, "")
// Output: 2022/03/08 11:47:45 there was an error: Error producing function exit status 1
