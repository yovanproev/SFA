package main

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

type CardComperator func(cOne Card, cTwo Card) int

var cardsDeck = Deck{
	Cards: []Card{{
		Number:   []string{"two", "three", "four", "five", "six", "seven", "eigth", "nine", "ten", "jack", "queen", "king", "ace"},
		Suit:     []string{"club", "diamond", "heart", "spade"},
		Strength: []int{1, 2},
	}},
}

// wasn't defined in the homework what the anonymous function should do exactly
var (
	anonymousFunc = func(cOne, cTwo Card) int {
		fmt.Println(cOne.Strength[0] * cTwo.Strength[1])
		return cOne.Strength[0] * cTwo.Strength[1]
	}
)

func main() {
	MaxCard(cardsDeck.Cards, compareCards)  // passing compareCards
	MaxCard(cardsDeck.Cards, anonymousFunc) // passing anonymous function
}

func MaxCard(cards []Card, comparatorFunc CardComperator) Card {

	var cardOne, cardSuit []int

	//	Turn strings to integer
	for idx := range cards[0].Number {
		cardOne = append(cardOne, idx+1)
	}

	for idx := range cards[0].Suit {
		cardSuit = append(cardSuit, idx+1)
	}

	var number, suit = findMaxCard(cardOne, cardSuit)

	fmt.Println("Strongest card is " + cards[0].Number[number-1] + " of " + cards[0].Suit[suit-1])

	comparatorFunc(cards[0], cards[0])

	return cards[0]
}

func findMaxCard(number, suit []int) (int, int) {
	maxNumber := number[0]
	maxSuit := suit[0]
	for _, value := range number {
		if value > maxNumber {
			maxNumber = value
		}
	}
	for _, value := range suit {
		if value > maxSuit {
			maxSuit = value
		}
	}
	return maxNumber, maxSuit
}

//No input on the console needed
// Output: "Strongest card is ace of spade"

func ConvertCardsToNumber(cardOneNumber, cardOneSuit, cardTwoNumber, cardTwoSuit string) []int {

	var cardOne, cardSuitOne, cardTwo, cardSuitTwo int

	for idx := range cardsDeck.Cards[0].Number {
		if strings.ToLower(cardOneNumber) == cardsDeck.Cards[0].Number[idx] {
			cardOne = idx + 2
		}
		if strings.ToLower(cardTwoNumber) == cardsDeck.Cards[0].Number[idx] {
			cardTwo = idx + 2
		}
	}

	for idx := range cardsDeck.Cards[0].Suit {
		if strings.ToLower(cardOneSuit) == cardsDeck.Cards[0].Suit[idx] {
			cardSuitOne = idx + 1
		}
		if strings.ToLower(cardTwoSuit) == cardsDeck.Cards[0].Suit[idx] {
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
	var sliceOfcards []int
	sliceOfcards = append(sliceOfcards, card1, card2)

	return sliceOfcards
}

func compareCards(cardOne Card, cardTwo Card) int {
	cardChoice1 := ConvertCardsToNumber(cardOne.Number[0], cardOne.Suit[0], cardTwo.Number[1], cardTwo.Suit[1])
	//	cardChoice2 := ConvertCardsToNumber(cardOne.Number[5], cardOne.Suit[2], cardTwo.Number[2], cardTwo.Suit[2])

	var result = 0
	var card1 int
	var card2 int

	card1 = cardChoice1[0]
	card2 = cardChoice1[1]

	if card1 < card2 {
		result = -1
	} else if card1 == 0 || card2 == 0 {
		return 00000
	} else if card1 == card2 {
		result = 0
	} else if card1 > card2 {
		result = 1
	}

	fmt.Println(result)
	return result
}

//No input on console needed
// Output:
//Strongest card is ace of spade
//-1
//Strongest card is ace of spade
//2
