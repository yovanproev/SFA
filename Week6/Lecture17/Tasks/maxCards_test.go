package compareCards

import "testing"

func TestMaxCard(t *testing.T) {

	var CardsDeckTest = Deck{
		Cards: []Card{{
			Number:   []string{"two", "three", "four", "five", "six", "seven", "eigth", "nine", "ten", "jack", "queen", "king", "ace"},
			Suit:     []string{"club", "diamond", "heart", "spade"},
			Strength: []int{4, 5, 6, 7, 8, 9},
		}},
	}

	// Act
	maxCard, maxSuit := findMaxCard(CardsDeckTest.Cards[0].Strength, CardsDeckTest.Cards[0].Strength)
	checkMaxCard := MaxCard(CardsDeckTest.Cards)

	// Assertion for FindMaxCard
	highestValue := CardsDeckTest.Cards[0].Strength[len(CardsDeckTest.Cards[0].Strength)-1]

	if maxCard == 0 || maxSuit == 0 || maxCard != highestValue || maxSuit != highestValue {
		t.Errorf("The value is %d, the suit is %d, please provide suitable card", maxCard, maxSuit)
	}

	// Assertion for MaxCard
	if checkMaxCard.Number == nil || checkMaxCard.Suit == nil {
		t.Errorf("The number is = %s and the suit %s; please provide a deck", checkMaxCard.Number, checkMaxCard.Suit)
	}
}

// Output
// $ go test ./... -cover
// ?       compareCards    [no test files]
// ok      compareCards/Tasks      (cached)        coverage: 83.7% of statements
