package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Type Deck represents the deck of cards which is []string.
type Deck []string

// Creates a new deck and return it.
func newDeck() Deck {
	var cards Deck
	cardSuites := []string{"Spades", "Diamond", "Hearts", "Sphere"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Print the contents of the deck
func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deals the card with handSize
func (d Deck) deal(handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

// Convert the deck to the string with comma seprated.
func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save Deck to the file
func (d Deck) saveDeckToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
