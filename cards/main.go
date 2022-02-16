package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()

	cardsToFile := newDeck()
	err := cardsToFile.saveToFile("my_cards")
	if err != nil {
		fmt.Println("Error saving file", err)
	}

	cardsFromFile := newDeckFromFile("my_cards")
	cardsFromFile.print()

	cardsShuffle := newDeck()
	cardsShuffle.shuffle()
	cardsShuffle.print()
}
