package main

import "fmt"

// create a new type 'deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// _ is go language thing for when you dont want to use the index variable of a for loop
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// receiver
// convention for receivers is to have the reference variable passed in be an abbreviation of the type
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// go can return multiple values from one function :0
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
