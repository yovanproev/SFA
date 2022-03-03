package main

import (
	cardgame "cardgame/Task1"
)

func main() {
	deck := cardgame.MakeDeck()
	deck.Shuffle()
	cardgame.Deal(deck, 10)
}
