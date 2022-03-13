package main

import (
	"cardgame/carddraw"
	"cardgame/cardgame"
	"log"
)

func main() {
	deck := cardgame.MakeDeck()
	deck.Shuffle()

	carddraw.DrawAllCards(&deck)
	if len(deck.Cards) == 0 {
		log.Fatal()
	}
}

// 52 cards drawn
// Output:
// First Draw  {Five Clubs}
// First Draw  {Jack Spades}
// First Draw  {Six Clubs}
// First Draw  {Jack Hearts}
// First Draw  {Six Spades}
// First Draw  {Queen Diamonds}
// First Draw  {Queen Hearts}
// First Draw  {Eigth Hearts}
// First Draw  {Nine Spades}
// First Draw  {King Hearts}
// First Draw  {Five Spades}
// First Draw  {Jack Diamonds}
// First Draw  {Ace Spades}
// First Draw  {Four Clubs}
// First Draw  {Four Hearts}
// First Draw  {Nine Diamonds}
// First Draw  {Seven Hearts}
// First Draw  {Ace Clubs}
// First Draw  {Five Hearts}
// First Draw  {Six Diamonds}
// First Draw  {King Diamonds}
// First Draw  {Seven Spades}
// First Draw  {Two Clubs}
// First Draw  {Ace Diamonds}
// First Draw  {Ace Hearts}
// First Draw  {Jack Clubs}
// First Draw  {Ten Clubs}
// First Draw  {Queen Spades}
// First Draw  {Ten Hearts}
// First Draw  {Seven Diamonds}
// First Draw  {Two Spades}
// First Draw  {Four Spades}
// First Draw  {Nine Clubs}
// First Draw  {Ten Diamonds}
// First Draw  {King Clubs}
// First Draw  {Seven Clubs}
// First Draw  {Eigth Clubs}
// First Draw  {Three Hearts}
// First Draw  {Three Diamonds}
// First Draw  {Six Hearts}
// First Draw  {Three Spades}
// First Draw  {Eigth Diamonds}
// First Draw  {Five Diamonds}
// First Draw  {Eigth Spades}
// First Draw  {Two Hearts}
// First Draw  {Two Diamonds}
// First Draw  {Three Clubs}
// First Draw  {Queen Clubs}
// First Draw  {Ten Spades}
// First Draw  {Nine Hearts}
// First Draw  {Four Diamonds}
// First Draw  {King Spades}
// no more cards in the deck -> <nil>
// 2022/03/06 18:48:36
// exit status 1
