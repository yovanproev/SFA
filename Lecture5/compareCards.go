package compareCards

import (
	"fmt"
	"strings"
)

type Card struct {
	Number []string
	Suit   []string
}

var CardsDeck = Card{
	Number: []string{"two", "three", "four", "five", "six", "seven", "eigth", "nine", "ten", "jack", "queen", "king", "ace"},
	Suit:   []string{"club", "diamond", "heart", "spade"},
}

func ConvertCardsToNumber(cardOneNumber, cardOneSuit, cardTwoNumber, cardTwoSuit string) (int, int) {

	var cardOne, cardSuitOne, cardTwo, cardSuitTwo int

	for idx := range CardsDeck.Number {
		if strings.ToLower(cardOneNumber) == CardsDeck.Number[idx] {
			cardOne = idx + 2
		}
		if strings.ToLower(cardTwoNumber) == CardsDeck.Number[idx] {
			cardTwo = idx + 2
		}
	}

	for idx := range CardsDeck.Suit {
		if strings.ToLower(cardOneSuit) == CardsDeck.Suit[idx] {
			cardSuitOne = idx + 1
		}
		if strings.ToLower(cardTwoSuit) == CardsDeck.Suit[idx] {
			cardSuitTwo = idx + 1
		}
	}

	var card1, card2 int

	if cardOne == 0 || cardTwo == 0 || cardSuitOne == 0 || cardSuitTwo == 0 {
		fmt.Println("The number/suit you inserted is invalid or invalid format, try again")
	} else if cardOne > 1 && cardTwo > 1 && cardOne < 15 && cardTwo < 15 &&
		cardSuitOne > 0 && cardSuitTwo > 0 && cardSuitOne <= 4 && cardSuitTwo <= 4 {

		card1 = cardOne + cardSuitOne
		card2 = cardTwo + cardSuitTwo
	} else {
		fmt.Println("The number/suit you inserted is invalid or invalid format, try again")
	}

	return card1, card2
}

func CompareCards(cardOne, cardTwo int) {

	if cardOne < cardTwo {
		fmt.Println(-1)
	} else if cardOne == 0 || cardTwo == 0 {
		return
	} else if cardOne == cardTwo {
		fmt.Println(0)
	} else if cardOne > cardTwo {
		fmt.Println(1)
	}
}

// Input: two spade three spade
// Output: -1

// Input: four heart three diamond
// Output: 1

// Input: three spade three spade
// Output: 0
