package main

import "fmt"

// import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println()
	remainingCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}