package main

import (
	compareCards "compareCards/lecture5"
	"fmt"
)

func main() {
	fmt.Print("Enter Card 1, Card 1 suit, Card 2, Card 2 suit: ")
	var card1, cardSuit1, card2, cardSuit2 string

	_, err := fmt.Scanf("%s %s %s %s", &card1, &cardSuit1, &card2, &cardSuit2)

	if err != nil {
		panic(err)
	}

	//Task 1
	convertedCard1, convertedCard2 := compareCards.ConvertCardsToNumber(card1, cardSuit1, card2, cardSuit2)
	compareCards.CompareCards(convertedCard1, convertedCard2)

	//Task 2
	compareCards.MaxCard(compareCards.CardsDeck)
}
