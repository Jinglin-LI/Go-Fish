package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	//card := "Ace of Spades"
	// var card = "test"
	// card = "Five of Diamonds"

	/*
		card := newCard()
		fmt.Println(card)
	*/
	/*
		cards := deck{"Ace of Diamonds", newCard()}
		cards = append(cards, "Six of Spades")
	*/

	// cards.print()

	/*
		for i, card := range cards { // i, card -> index, value
			fmt.Println(i, card)
		}
	*/

	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// cards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	cards := newDeck()
	fmt.Println(cards.toString())
}

func newCard() string {
	return "Five of Diamonds"
}
