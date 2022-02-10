package main

import "fmt"

func main() {
	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	cards := newDeck()

	err := cards.saveToFile("my_cards")
	if err != nil {
		fmt.Println("Error saving file", err)
	}
}
