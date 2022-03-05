package main

import (
	"cardgame/cardgame"
)

func main() {
	deck := cardgame.MakeDeck()
	deck.Shuffle()

	ml := cardgame.Card2{}

	node1 := &cardgame.Deck2{
		Value: deck.Value,
	}
	node2 := &cardgame.Deck2{
		Value: deck.Value,
	}
	node3 := &cardgame.Deck2{
		Value: deck.Value,
	}
	node4 := &cardgame.Deck2{
		Value: deck.Value,
	}
	ml.Deal2(node1)
	ml.Deal2(node2)
	ml.Deal2(node3)
	ml.Deal2(node4)

	//carddraw.DrawAllCards(ml.Deal2())

	cardgame.ToSlice(ml)
}
