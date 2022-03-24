package compareCards

import (
	"fmt"
	"strings"
)

type Deck struct {
	Cards []Card
}

type Card struct {
	Number   []string
	Suit     []string
	Strength []int
}

var CardsDeck = Deck{
	Cards: []Card{{
		Number: []string{"two", "three", "four", "five", "six", "seven", "eigth", "nine", "ten", "jack", "queen", "king", "ace"},
		Suit:   []string{"club", "diamond", "heart", "spade"},
	}},
}

func ConvertCardsToNumber(cardOneNumber, cardOneSuit, cardTwoNumber, cardTwoSuit string) Card {

	var cardOne, cardSuitOne, cardTwo, cardSuitTwo int

	for idx := range CardsDeck.Cards[0].Number {
		if strings.ToLower(cardOneNumber) == CardsDeck.Cards[0].Number[idx] {
			cardOne = idx + 2
		}
		if strings.ToLower(cardTwoNumber) == CardsDeck.Cards[0].Number[idx] {
			cardTwo = idx + 2
		}
	}

	for idx := range CardsDeck.Cards[0].Suit {
		if strings.ToLower(cardOneSuit) == CardsDeck.Cards[0].Suit[idx] {
			cardSuitOne = idx + 1
		}
		if strings.ToLower(cardTwoSuit) == CardsDeck.Cards[0].Suit[idx] {
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
	CardsDeck.Cards[0].Strength = append(CardsDeck.Cards[0].Strength, card1, card2)

	return CardsDeck.Cards[0]
}

func CompareCards(cardOne Card, cardTwo Card) int {
	var result int
	var card1 = CardsDeck.Cards[0].Strength[0]
	var card2 = CardsDeck.Cards[0].Strength[1]
	if card1 < card2 {
		result = -1
	} else if card1 == 0 || card2 == 0 {
		return 00000
	} else if card1 == card2 {
		result = 0
	} else if card1 > card2 {
		result = 1
	}
	return result
}

// Input on console: two spade three spade
// Output: -1

// Input on console: four heart three diamond
// Output: 1

// Input on console: three spade three spade
// Output: 0
