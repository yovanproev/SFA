package cardgame

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck is a collection of cards
type Deck struct {
	Cards []Card
}

// Card is what makes up a deck
type Card struct {
	Value string
	Suite string
}

func MakeDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eigth", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards.Cards = append(cards.Cards, Card{Value: value + " of", Suite: suite})
		}
	}

	return cards
}

func (d *Deck) Deal() *Card {

	for i, card := range d.Cards {
		fmt.Println(i+1, card)
	}
	fmt.Println("The deck is empty")
	return &d.Cards[0]
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d.Cards {
		newPosition := r.Intn(len(d.Cards) - 1)

		d.Cards[i], d.Cards[newPosition] = d.Cards[newPosition], d.Cards[i]
	}
}
