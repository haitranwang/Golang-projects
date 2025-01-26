package main

func main() {
	cards := desk{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
