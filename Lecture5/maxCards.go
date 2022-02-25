package compareCards

import "fmt"

func MaxCard(cards Card) (string, string) {
	var cardOne, cardSuit []int
	// Turn strings to integer
	for idx := range CardsDeck.Number {
		cardOne = append(cardOne, idx+2)
	}

	for idx := range CardsDeck.Suit {
		cardSuit = append(cardSuit, idx+1)
	}

	var number, suit = findMaxCard(cardOne, cardSuit)

	// Turn integer to string again
	var numberBackToString, suitBackToString string

	for idx, v := range CardsDeck.Number {
		if idx+2 == number {
			numberBackToString = v
		}
	}

	for idx, v := range CardsDeck.Suit {
		if idx+1 == suit {
			suitBackToString = v
		}
	}
	fmt.Println("Strongest card is " + numberBackToString + " of " + suitBackToString)

	return numberBackToString, suitBackToString
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

//No input on the console needed
// Output: "Strongest card is ace of spade"
