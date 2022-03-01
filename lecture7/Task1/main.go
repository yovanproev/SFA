// Example only with slices
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck []string

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 20)
	cards.shuffle()
	hand.print()
	remainingCards.print()
}

func newDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eigth", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	if handSize == 52 {
		fmt.Println("No more cards in the deck")
	}
	return d[:handSize], d[handSize:]
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
