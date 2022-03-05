package carddraw

import (
	"cardgame/cardgame"
)

type Dealer interface {
	Deal2() *cardgame.Card2
}

func DrawAllCards(dealer Dealer) []cardgame.Card2 {

	dealer.Deal2()

	return []cardgame.Card2{}
}
