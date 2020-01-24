package main

// Array are fixed length and slice can shrink/grow

func main() {
	cards := newDeck()
	cards.cardRandomize()
	cards.print()

	even()
}
