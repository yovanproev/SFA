package compareCards

import "fmt"

//func maxCard(cards []Card) Card { // use compareCards here to find the maximum ... }
func MaxCard(cards []Card) Card {

	var cardOne, cardSuit []int

	//	Turn strings to integer
	for idx := range cards[0].Number {
		cardOne = append(cardOne, idx+1)
	}

	for idx := range cards[0].Suit {
		cardSuit = append(cardSuit, idx+1)
	}

	var number, suit = findMaxCard(cardOne, cardSuit)

	fmt.Println("Strongest card is " + cards[0].Number[number-1] + " of " + cards[0].Suit[suit-1])
	return cards[0]
}

func findMaxCard(number, suit []int) (int, int) {
	maxNumber := number[0]
	maxSuit := suit[0]
	for _, value := range number {
		if value > maxNumber {
			maxNumber = value
		}
	}
	for _, value := range suit {
		if value > maxSuit {
			maxSuit = value
		}
	}
	return maxNumber, maxSuit
}

//No input on the console needed. From a full deck of cards, provided from compareCards file struct.
// Output: "Strongest card is ace of spade"
