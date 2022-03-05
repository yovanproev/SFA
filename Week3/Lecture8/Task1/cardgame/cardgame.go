package cardgame

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Value string
	Suite string
	prev  *Card
	next  *Card
	key   interface{}
}

type Deck struct {
	Cards []Card
	head  *Card
	tail  *Card
}

// func MakeDeck() Deck {
// 	cards := Deck{}

// 	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
// 	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eigth", "Nine", "Ten", "Jack", "Queen", "King"}

// 	for _, suite := range cardSuits {
// 		for _, value := range cardValues {
// 			cards.Cards = append(cards.Cards, Card{Value: value + " of", Suite: suite})
// 		}
// 	}
// 	//fmt.Println(cards.Cards[0].Value) //Ace of
// 	return cards
// }

func (d *Deck) Deal() *Card {
	list := &Card{
		next: d.head,
		key:  d.Cards[0].key,
	}
	if d.head != nil {
		d.head.prev = list
	}
	d.head = list

	l := d.head
	for l.next != nil {
		l = l.next
	}
	d.tail = l

	//	fmt.Println(d.Cards)
	//fmt.Println(d.Cards)
	return &d.Cards[0]
}

func (l *Deck) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func (d Deck2) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d.Value {
		newPosition := r.Intn(len(d.Value) - 1)

		d.Value[i], d.Value[newPosition] = d.Value[newPosition], d.Value[i]
	}
}

type Deck2 struct {
	Value    []string
	PrevItem *Deck2
}

type Card2 struct {
	LastItem *Deck2
	length   int
}

func MakeDeck() Deck2 {
	var cards = Deck2{}
	// var cards []string

	var cardSuits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eigth", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards.Value = append(cards.Value, value+" of "+suite)
		}
	}
	//fmt.Println(cards) //Ace of
	return cards
}

func (l *Card2) Deal2(i *Deck2) {
	second := l.LastItem
	l.LastItem = i
	l.LastItem.PrevItem = second
	l.length++
}

func ToSlice(ml Card2) []string {
	var slice []string
	var restOfDeck []string

	//toPrint := ml.LastItem

	for i := 0; i < ml.length; i++ {
		//fmt.Println(ml.length)
		restOfDeck = nil
		slice = append(slice, ml.LastItem.Value[i])
		restOfDeck = append(restOfDeck, ml.LastItem.Value[ml.length:52]...)
	}

	if ml.LastItem == nil {
		fmt.Println("No Deck initialized")
	} else if restOfDeck == nil {
		fmt.Println("Deck is empty")
	} else {
		fmt.Println("First Draw", slice)
		fmt.Println("Rest of the Deck", restOfDeck)
	}
	return slice
}
