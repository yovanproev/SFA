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

//10 randomly drawn cards.
// Output:
// First Draw  {Eigth Spades}
// First Draw  {Ace Hearts}
// First Draw  {Queen Hearts}
// First Draw  {Six Diamonds}
// First Draw  {Four Hearts}
// First Draw  {Seven Hearts}
// First Draw  {Nine Clubs}
// First Draw  {Jack Clubs}
// First Draw  {Two Spades}
// First Draw  {Six Hearts}
// Rest of Deck:  &{[{Five Hearts} {Ten Diamonds} {King Spades} {King Diamonds} {Three Diamonds} {Three Clubs} {Four Diamonds} {Eigth Clubs} {Two Hearts} {Nine Diamonds} {Nine Hearts} {Jack Hearts} {Two Diamonds} {Nine Spades} {Five Diamonds} {Ten Spades} {Queen Clubs} {Seven Diamonds} {Queen Diamonds} {Seven Clubs} {Eigth Diamonds} {Five Clubs} {Ace Diamonds} {Queen Spades} {Jack Spades} {Three Spades} {King Hearts} {Five Spades} {Six Clubs} {Ten Clubs} {Ten Hearts} {Two Clubs} {Ace Spades} {Jack Diamonds} {Three Hearts} {Four Spades} {King Clubs} {Eigth Hearts} {Four Clubs}
// {Six Spades} {Seven Spades} {Ace Clubs}]}
