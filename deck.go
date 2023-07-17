package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// Read Deck From File
func readDeckFromFile(fileName string) Deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := string(bs)
	return Deck(strings.Split(s, ","))
}

func (d Deck) shuffle() Deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}

	return d
}
