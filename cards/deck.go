package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of "deck"
// which is a slice of strings.
type deck []string

// Create new deck which contain 16 cards.
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print all cards in deck
// each on newline with index and value.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Make a "hand" of cards
// from existing deck.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Convert deck to string representation.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save existing deck to specified file.
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Create a deck from specified file.
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Shuffle existing deck.
func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())

	for i := range d {
		np := rand.Intn(len(d))

		d[i], d[np] = d[np], d[i]
	}
}
