package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create type deck
type deck []string

// This function for create new deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// This funcation is for for print deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// This function is for dealing the cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// This function is for returning string from slice
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// This function save the deck in to the file locally
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// This function return deck from saved file locally
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// This function will shuffle the deck
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // Genrate seed using unix time for random new function
	r := rand.New(source)                           // Genrate new random number using above seed

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] //swaping the values in slice
	}
}
