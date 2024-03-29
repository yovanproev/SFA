package cardgame

import (
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
