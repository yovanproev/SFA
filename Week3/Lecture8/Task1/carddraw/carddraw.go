package carddraw

import (
	"cardgame/cardgame"
	"fmt"
)

type Dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(dealer Dealer) []cardgame.Card {

	for i := 0; i < 10; i++ {
		pointerToCard := *dealer.Deal()
		fmt.Println("First Draw ", pointerToCard)
	}
	fmt.Println("Rest of Deck: ", dealer)

	return []cardgame.Card{}
}
