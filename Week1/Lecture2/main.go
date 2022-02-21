package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter Card 1, Card 1 suit, Card 2, Card 2 suit: ")
	var card1, card2 int
	var cardSuit1, cardSuit2 string

	_, err := fmt.Scanf("%d %s %d %s", &card1, &cardSuit1, &card2, &cardSuit2)

	if err != nil {
		panic(err)
	}

	compareCards(card1, cardSuit1, card2, cardSuit2)
}

func convertCardSuits(cardOneSuit, cardTwoSuit string) (int, int) {
	cardSuits := []string{"club", "diamond", "heart", "spade"}
	var number1, number2 int

	for i := range cardSuits {
		if strings.ToLower(cardOneSuit) == cardSuits[i] {
			number1 = i + 1
		}
		if strings.ToLower(cardTwoSuit) == cardSuits[i] {
			number2 = i + 1
		}
	}
	return number1, number2
}

func compareCards(cardOneVal int, cardOneSuit string, cardTwoVal int, cardTwoSuit string) {
	var cardSuit1, cardSuit2 = convertCardSuits(cardOneSuit, cardTwoSuit)

	cardOne := cardOneVal + cardSuit1
	cardTwo := cardTwoVal + cardSuit2

	if cardOneVal > 1 && cardTwoVal > 1 && cardOneVal < 14 && cardTwoVal < 14 &&
		cardSuit1 > 0 && cardSuit2 > 0 && cardSuit1 <= 4 && cardSuit2 <= 4 {
		if cardOne < cardTwo {
			fmt.Println(-1)
		} else if cardOne == cardTwo {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	} else {
		fmt.Println("Invalid number or suit, try again")
	}
}
