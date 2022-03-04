package main

import (
	"cardgame/carddraw"
	"cardgame/cardgame"
)

func main() {
	deck := cardgame.MakeDeck()
	deck.Shuffle()
	carddraw.DrawAllCards(&deck)
}
