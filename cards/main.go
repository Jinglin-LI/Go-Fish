package main

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

	cards := newDeck()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
