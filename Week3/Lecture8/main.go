package main

import (
	cardgame "cardgame/Task1"
)

func main() {
	deck := cardgame.MakeDeck()
	deck.Shuffle()
	deck.Deal(10)
}
