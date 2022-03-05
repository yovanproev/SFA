package cardgame

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	Value    []string
	PrevItem *Deck
}

type Card struct {
	LastItem *Deck
	length   int
}

func MakeDeck() Deck {
	var cards = Deck{}

	var cardSuits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eigth", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards.Value = append(cards.Value, value+" of "+suite)
		}
	}
	return cards
}

func (c *Card) Deal(d *Deck) {
	second := c.LastItem
	c.LastItem = d
	c.LastItem.PrevItem = second
	c.length++
}

func ToSlice(c Card) []string {
	var slice []string
	var restOfDeck []string

	for i := 0; i < c.length; i++ {
		restOfDeck = nil
		slice = append(slice, c.LastItem.Value[i])
		restOfDeck = append(restOfDeck, c.LastItem.Value[c.length:52]...)
	}

	if c.LastItem == nil {
		fmt.Println("No Deck initialized")
	} else if restOfDeck == nil {
		fmt.Println("Deck is empty")
	} else {
		fmt.Println("First Draw", slice)
		fmt.Println("Rest of the Deck", restOfDeck)
	}
	return slice
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d.Value {
		newPosition := r.Intn(len(d.Value) - 1)

		d.Value[i], d.Value[newPosition] = d.Value[newPosition], d.Value[i]
	}
}
