package carddraw

import (
	"cardgame/cardgame"
	"fmt"
)

type Dealer interface {
	Deal() (*cardgame.Card, error)
	Done() bool
}

func DrawAllCards(dealer Dealer) []cardgame.Card {

	for i := 0; i < 52; i++ {
		pointerToCard, _ := dealer.Deal()
		fmt.Println("First Draw ", *pointerToCard)
	}

	if dealer.Done() {
		pointerToCard, err := dealer.Deal()
		fmt.Println(err.Error(), pointerToCard)
	} else {
		fmt.Println("Rest of Deck: ", dealer)
	}

	return []cardgame.Card{}
}
