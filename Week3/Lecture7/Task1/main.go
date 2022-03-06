package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

type Card struct {
	Value string
	Suite string
}

func main() {
	deck := MakeDeck()
	deck.Shuffle()

	for i := 0; i < 52; i++ {
		pointerToCard := *deck.Deal()
		fmt.Println("First Draw ", pointerToCard)

	}
	fmt.Println("Rest of Deck: ", deck.Cards)
}

func MakeDeck() Deck {
	var deck = Deck{}

	var cardSuits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eigth", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			var card = Card{Value: value, Suite: suite}

			deck.Cards = append(deck.Cards, card)
		}
	}

	return deck
}

func (d *Deck) Deal() *Card {

	if len(d.Cards) == 0 {
		return nil
	}
	firstCard := d.Cards[0]
	d.Cards = d.Cards[1:len(d.Cards)]

	return &firstCard
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d.Cards {
		newPosition := r.Intn(len(d.Cards) - 1)

		d.Cards[i], d.Cards[newPosition] = d.Cards[newPosition], d.Cards[i]
	}
}

// Random drawing cards. 5 cards drawn from deck.
//Output:
//First Draw  {Nine Hearts}
//First Draw  {Nine Clubs}
//First Draw  {King Diamonds}
//First Draw  {Seven Spades}
//First Draw  {Ace Diamonds}
//Rest of Deck:  [{Ten Hearts} {Nine Spades} {Jack Clubs} {Six Diamonds} {Queen Clubs} {Five Diamonds} {Eigth Spades} {Six Spades} {Five Clubs} {Four Hearts} {Eigth Hearts} {Six Clubs} {Queen Hearts} {Jack Hearts} {Three Hearts} {King Spades} {Ace Hearts} {Nine Diamonds} {Two Spades} {King Hearts} {Four Diamonds} {Eigth Clubs} {Jack Spades} {Two Clubs} {Ten Clubs} {Seven Diamonds} {Five Spades} {Ace Clubs} {Four Clubs} {Seven Clubs}
//{Six Hearts} {Three Diamonds} {Two Diamonds} {Jack Diamonds} {King Clubs} {Eigth Diamonds} {Ten Diamonds} {Three Clubs} {Queen Spades} {Ace Spades} {Two Hearts} {Queen Diamonds} {Three Spades} {Four Spades} {Five Hearts} {Ten Spades} {Seven Hearts}]

// Random drawing cards. 52 cards drawn from deck.
//Output:
// First Draw  {Ace Hearts}
// First Draw  {Three Spades}
// First Draw  {Six Hearts}
// First Draw  {Six Diamonds}
// First Draw  {Two Hearts}
// First Draw  {Eigth Hearts}
// First Draw  {Queen Clubs}
// First Draw  {Ace Diamonds}
// First Draw  {Ten Diamonds}
// First Draw  {Nine Spades}
// First Draw  {Six Spades}
// First Draw  {King Spades}
// First Draw  {Two Diamonds}
// First Draw  {Jack Hearts}
// First Draw  {Ten Spades}
// First Draw  {Four Spades}
// First Draw  {Three Clubs}
// First Draw  {Eigth Spades}
// First Draw  {King Clubs}
// First Draw  {Queen Hearts}
// First Draw  {Queen Diamonds}
// First Draw  {Seven Diamonds}
// First Draw  {Ten Hearts}
// First Draw  {Five Spades}
// First Draw  {Two Spades}
// First Draw  {Nine Clubs}
// First Draw  {Eigth Clubs}
// First Draw  {Five Clubs}
// First Draw  {Seven Spades}
// First Draw  {Ace Clubs}
// First Draw  {King Diamonds}
// First Draw  {Five Hearts}
// First Draw  {Five Diamonds}
// First Draw  {Four Hearts}
// First Draw  {Seven Clubs}
// First Draw  {Nine Diamonds}
// First Draw  {Four Clubs}
// First Draw  {Jack Spades}
// First Draw  {Seven Hearts}
// First Draw  {Four Diamonds}
// First Draw  {Nine Hearts}
// First Draw  {Ace Spades}
// First Draw  {Jack Clubs}
// First Draw  {Two Clubs}
// First Draw  {Queen Spades}
// First Draw  {Ten Clubs}
// First Draw  {Eigth Diamonds}
// First Draw  {Six Clubs}
// First Draw  {King Hearts}
// First Draw  {Three Hearts}
// First Draw  {Three Diamonds}
// First Draw  {Jack Diamonds}
// Rest of Deck:  []
