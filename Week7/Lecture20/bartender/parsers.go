package cocktails

import (
	"regexp"
	"strings"
)

func turnInputToRunes(s string) ([]rune, string) {
	var turnInputToRunes []rune

	cutReaderNewLineChar := cutReaderNewLineChar(s)
	everyFirstLetterCapital := strings.Title(strings.ToLower(cutReaderNewLineChar))

	if smallLetterAfterQuotationMarks(cutReaderNewLineChar) != nil {
		modify := strings.Replace(everyFirstLetterCapital, "'S", string(smallLetterAfterQuotationMarks(cutReaderNewLineChar)[0][0]), -1)
		turnInputToRunes = adjustTheRuneSlice(modify)
	} else {
		turnInputToRunes = adjustTheRuneSlice(everyFirstLetterCapital)
	}

	return turnInputToRunes, everyFirstLetterCapital
}

func smallLetterAfterQuotationMarks(str string) [][][]byte {
	re := regexp.MustCompile(`'(.?)`)
	smallLetterAfterQuotationMarks := re.FindAllSubmatch([]byte(str), -1)

	return smallLetterAfterQuotationMarks
}

func adjustTheRuneSlice(str string) []rune {
	turnInputToRunes := []rune(str)
	return turnInputToRunes
}

func cutReaderNewLineChar(s string) string {
	consoleInput := strings.Replace(s, "\n", "", -1)
	consoleInput = strings.Replace(consoleInput, "\r", "", -1)

	return consoleInput
}

func runesToStrings(s string, drink string) string {
	var runesToStrings string

	everyFirstLetterCapital := strings.Title(strings.ToLower(s))

	if smallLetterAfterQuotationMarks(s) != nil {
		modify := strings.Replace(everyFirstLetterCapital, "'S", string(smallLetterAfterQuotationMarks(s)[0][0]), -1)
		adjustTheRuneSlice := adjustTheRuneSlice(modify)[0:len(drink)]
		runesToStrings = string(adjustTheRuneSlice)
	} else {
		adjustTheRuneSlice := adjustTheRuneSlice(everyFirstLetterCapital)[0:len(drink)]
		runesToStrings = string(adjustTheRuneSlice)
	}

	return runesToStrings
}
