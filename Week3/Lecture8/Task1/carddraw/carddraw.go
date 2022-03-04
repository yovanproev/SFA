package carddraw

import (
	"cardgame/cardgame"
)

type Dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(dealer Dealer) []cardgame.Card {
	dealer.Deal()

	return cardgame.MakeDeck().Cards
}
