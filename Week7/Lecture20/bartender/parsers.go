package cocktails

import (
	"regexp"
	"strings"
)

func TurnInputToRunes(s string) ([]rune, string) {
	var turnInputToRunes []rune

	modifyInputToAcceptedFormat := strings.Replace(s, "\n", "", -1)
	modifyInputToAcceptedFormat = strings.Replace(modifyInputToAcceptedFormat, "\r", "", -1)
	everyFirstLetterCapital := strings.Title(strings.ToLower(string(modifyInputToAcceptedFormat)))

	if smallLetterAfterQuotationMarks(modifyInputToAcceptedFormat) != nil {
		modify := strings.Replace(everyFirstLetterCapital, "'S", string(smallLetterAfterQuotationMarks(modifyInputToAcceptedFormat)[0][0]), -1)
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

func TurnInputToURLFormat(s string) string {
	// drinks with Quotation mark
	modifyToURLAcceptableFormat := strings.Replace(s, "'", "%27", -1)

	everyFirstLetterCapital := strings.Title(strings.ToLower(modifyToURLAcceptableFormat))
	modifyToURLAcceptableFormat = strings.Replace(everyFirstLetterCapital, " ", "%20", -1)
	modifyToURLAcceptableFormat = strings.Replace(modifyToURLAcceptableFormat, "\n", "", -1)
	modifyToURLAcceptableFormat = strings.Replace(modifyToURLAcceptableFormat, "\r", "", -1)

	return modifyToURLAcceptableFormat
}

func RunesToStrings(s string, drink string) string {
	var runesToStrings string

	if smallLetterAfterQuotationMarks(s) != nil {
		everyFirstLetterCapital := strings.Title(strings.ToLower(s))
		modify := strings.Replace(everyFirstLetterCapital, "'S", string(smallLetterAfterQuotationMarks(s)[0][0]), -1)
		adjustTheRuneSlice := adjustTheRuneSlice(modify)[0:len(drink)]
		runesToStrings = string(adjustTheRuneSlice)
	} else {
		everyFirstLetterCapital := strings.Title(strings.ToLower(s))
		adjustTheRuneSlice := adjustTheRuneSlice(everyFirstLetterCapital)[0:len(drink)]
		runesToStrings = string(adjustTheRuneSlice)
	}

	return runesToStrings
}
