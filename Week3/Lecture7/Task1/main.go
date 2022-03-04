package main

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

func main() {
	deck := makeDeck()
	deck.shuffle()
	deck.deal(10)
}

func makeDeck() Deck {
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

func (d Deck) deal(handSize int) {
	if handSize == 52 {
		fmt.Println("No more cards in the deck")
	}
	for i := 0; i < handSize; i++ {
		fmt.Print(d.Cards[i])
	}
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d.Cards {
		newPosition := r.Intn(len(d.Cards) - 1)

		d.Cards[i], d.Cards[newPosition] = d.Cards[newPosition], d.Cards[i]
	}
}

// deck.deal(52)
//Output: No more cards in the deck
//{Diamonds of Six}{Hearts of Three}{Spades of Eigth}{Clubs of Four}{Hearts of Two}{Spades of Ace}{Spades of Two}{Clubs of Ace}{Diamonds of Seven}{Spades of Ten}{Clubs of Three}{Hearts of Seven}{Hearts of Jack}{Hearts of Ace}{Diamonds of Three}{Hearts of Four}{Diamonds of Four}{Hearts of Six}{Clubs of Six}{Clubs of Seven}{Diamonds of Jack}{Hearts of Five}{Spades of Three}{Hearts of Queen}{Diamonds of Two}{Clubs of Queen}{Spades of Nine}{Hearts of King}{Diamonds of Ace}{Diamonds of King}{Clubs of Eigth}{Clubs of Five}{Hearts of Nine}{Clubs of Nine}{Hearts of Ten}{Hearts of Eigth}{Diamonds of Nine}{Clubs of Ten}{Spades of Five}{Spades of King}{Clubs of King}{Diamonds of Ten}{Diamonds of Five}{Spades of Jack}{Spades of Queen}{Spades of Seven}{Spades of Four}{Diamonds of Queen}{Diamonds of Eigth}{Clubs of Jack}{Clubs of Two}{Spades of Six}

// deck.deal(10)
//Output:{Spades of Four}{Spades of Two}{Hearts of Seven}{Diamonds of Four}{Hearts of Two}{Clubs of King}{Spades of Queen}{Spades of King}{Spades of Seven}{Clubs of Ace}
