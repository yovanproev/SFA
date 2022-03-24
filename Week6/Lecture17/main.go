package main

import (
	cardgame "cardgame/cardgame"
	"fmt"
)

func main() {
	fmt.Print("Enter Card 1, Card 1 suit, Card 2, Card 2 suit: ")
	var card1, cardSuit1, card2, cardSuit2 string

	_, err := fmt.Scanf("%s %s %s %s\n", &card1, &cardSuit1, &card2, &cardSuit2)

	if err != nil {
		fmt.Println(err)
	}

	//Task 1
	Cards := cardgame.ConvertCardsToNumber(card1, cardSuit1, card2, cardSuit2)
	fmt.Println(cardgame.CompareCards(Cards, Cards))

	// Task 2
	strongestCard := cardgame.MaxCard(cardgame.CardsDeck.Cards)
	fmt.Println("Strongest card is " + strongestCard.Number[0] + " of " + strongestCard.Suit[0])
}
