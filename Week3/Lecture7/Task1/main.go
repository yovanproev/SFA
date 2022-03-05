package main

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

func main() {
	deck := MakeDeck()
	deck.Shuffle()

	d := Card{}

	node1 := &Deck{
		Value: deck.Value,
	}
	node2 := &Deck{
		Value: deck.Value,
	}
	node3 := &Deck{
		Value: deck.Value,
	}
	node4 := &Deck{
		Value: deck.Value,
	}
	d.Deal(node1)
	d.Deal(node2)
	d.Deal(node3)
	d.Deal(node4)
	ToSlice(d)
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

// Random drawing cards. 4 cards drawn from deck.
//Output:
//First Draw [Five of Clubs Jack of Hearts Queen of Clubs Three of Hearts]
//Rest of the Deck [Five of Diamonds Two of Hearts Four of Spades Eigth of Hearts Two of Diamonds Three of Clubs Queen of Spades Ten of Hearts Five of Spades Three of Spades Ace of Diamonds Six of Diamonds Eigth of Spades Queen of Diamonds Four of Hearts Five of Hearts Nine of Clubs King of
//Clubs Six of Hearts Four of Diamonds Seven of Hearts Six of Clubs Three of Diamonds Ten of Diamonds Ace of Hearts Four of Clubs Seven of Spades Nine of Diamonds Ten of Clubs Eigth of Diamonds Eigth of Clubs Six of Spades Ace of Spades Nine of Hearts King of Hearts Jack of Diamonds King of
//Spades Jack of Clubs Seven of Diamonds Queen of Hearts Seven of Clubs Two of Clubs Nine of Spades Ten of Spades King of Diamonds Jack of Spades Ace of Clubs Two of Spades]
