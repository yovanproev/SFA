package carddraw

import cardgame "cardgame/Task1"

type Dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(dealer Dealer) {

}
