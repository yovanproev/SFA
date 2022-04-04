package cocktails

import (
	"regexp"
	"strings"
)

func TurnInputToRunes(c CocktailBartender) ([]rune, string) {
	var turnInputToRunes []rune

	modifyInputToAcceptedFormat := strings.Replace(c.UserInput, "\n", "", -1)
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

func TurnInputToURLFormat(c CocktailBartender) string {
	// drinks with Quotation mark
	modifyToURLAcceptableFormat := strings.Replace(c.UserInput, "'", "%27", -1)

	everyFirstLetterCapital := strings.Title(strings.ToLower(modifyToURLAcceptableFormat))
	modifyToURLAcceptableFormat = strings.Replace(everyFirstLetterCapital, " ", "%20", -1)
	modifyToURLAcceptableFormat = strings.Replace(modifyToURLAcceptableFormat, "\n", "", -1)
	modifyToURLAcceptableFormat = strings.Replace(modifyToURLAcceptableFormat, "\r", "", -1)

	return modifyToURLAcceptableFormat
}

func RunesToStrings(c CocktailBartender, drink string) string {
	var runesToStrings string

	if smallLetterAfterQuotationMarks(c.UserInput) != nil {
		everyFirstLetterCapital := strings.Title(strings.ToLower(c.UserInput))
		modify := strings.Replace(everyFirstLetterCapital, "'S", string(smallLetterAfterQuotationMarks(c.UserInput)[0][0]), -1)
		adjustTheRuneSlice := adjustTheRuneSlice(modify)[0:len(drink)]
		runesToStrings = string(adjustTheRuneSlice)
	} else {
		everyFirstLetterCapital := strings.Title(strings.ToLower(c.UserInput))
		adjustTheRuneSlice := adjustTheRuneSlice(everyFirstLetterCapital)[0:len(drink)]
		runesToStrings = string(adjustTheRuneSlice)
	}

	return runesToStrings
}
