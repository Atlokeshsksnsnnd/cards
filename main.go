package main

const name = "lojesh"

func main() {

	cards := newDeck()
	cards.saveDeckToFile("cards.txt")

}
