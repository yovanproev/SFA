package cardgame

import (
	"math/rand"
	"testing"
	"time"
)

func TestCompareCards(t *testing.T) {
	rand.Seed(time.Now().Unix())

	var CardsDeckTest = Deck{
		Cards: []Card{{
			Number: []string{"two", "three", "four", "five", "six", "seven", "eigth", "nine", "ten", "jack", "queen", "king", "ace"},
			Suit:   []string{"club", "diamond", "heart", "spade"},
		}},
	}

	randomizeCard1 := rand.Intn(13)
	randomizeSuit1 := rand.Intn(4)
	randomizeCard2 := rand.Intn(13)
	randomizeSuit2 := rand.Intn(4)

	card1 := CardsDeckTest.Cards[0].Number[randomizeCard1]
	cardSuit1 := CardsDeckTest.Cards[0].Suit[randomizeSuit1]
	card2 := CardsDeckTest.Cards[0].Number[randomizeCard2]
	cardSuit2 := CardsDeckTest.Cards[0].Suit[randomizeSuit2]

	// Act
	cards := ConvertCardsToNumber(card1, cardSuit1, card2, cardSuit2)
	compare := CompareCards(cards, cards)

	// Assertion for ConvertCardsToNumber func
	for idx, stringValue := range cards.Number { // range over slice of strings of numbers
		if stringValue == "" || stringValue != CardsDeckTest.Cards[0].Number[idx] {
			t.Error("No number for comparison provided ", stringValue)
		}
	}

	for idx, stringValue := range cards.Suit { // range over slice of strings of suits
		if stringValue == "" || stringValue != CardsDeckTest.Cards[0].Suit[idx] {
			t.Error("No suit for comparison provided ", stringValue)
		}
	}

	for _, strengthOfCard := range cards.Strength { // range over slice of ints
		if strengthOfCard == 0 {
			t.Error("No strength of Card for comparison provided ", strengthOfCard)
		}
	}

	// Assertion for CompareCards func
	if compare > 1 || compare < -1 {
		t.Errorf("Result is = %d; want 0, 1 or -1", compare)
	}

}

// go test ./cardgame -cover
// Output
// ok      cardgame/cardgame       0.086s  coverage: 86.0% of statements
