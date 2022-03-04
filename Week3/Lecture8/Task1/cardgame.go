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
	Suite string
	Value string
}

func MakeDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eigth", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards.Cards = append(cards.Cards, Card{Suite: suite + " of", Value: value})
		}
	}

	return cards
}

func (d Deck) Deal(handSize int) {
	if handSize == 52 {
		fmt.Println("No more cards in the deck")
	}
	for i := 0; i < handSize; i++ {
		fmt.Println(d.Cards[i])
	}
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d.Cards {
		newPosition := r.Intn(len(d.Cards) - 1)

		d.Cards[i], d.Cards[newPosition] = d.Cards[newPosition], d.Cards[i]
	}
}
