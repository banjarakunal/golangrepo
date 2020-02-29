package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("Shuffle the deck:")
	cards.shuffle()
	cards.print()
	cards.saveToFile("my_cards")
	fmt.Println(cards.toString())
	cards = newDeckFromFile("my_cards")
	fmt.Println("Reading from file")
	cards.print()
	fmt.Println("...............")
	cards.print()
	hand, remainingHand := deal(cards, 5)
	cards.print()
	hand.print()
	remainingHand.print()

}
