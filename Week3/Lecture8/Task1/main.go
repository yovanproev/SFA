package main

import (
	"cardgame/cardgame"
)

func main() {
	deck := cardgame.MakeDeck()
	deck.Shuffle()

	d := cardgame.Card{}

	node1 := &cardgame.Deck{
		Value: deck.Value,
	}
	// node2 := &cardgame.Deck{
	// 	Value: deck.Value,
	// }
	// node3 := &cardgame.Deck{
	// 	Value: deck.Value,
	// }
	// node4 := &cardgame.Deck{
	// 	Value: deck.Value,
	// }
	d.Deal(node1)
	// d.Deal(node2)
	// d.Deal(node3)
	// d.Deal(node4)

	//carddraw.DrawAllCards(d)

	cardgame.ToSlice(d)
}
