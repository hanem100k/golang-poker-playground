package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}

	for _, cv := range cardValues {
		for _, cs := range cardSuits {
			cards = append(cards, cv+" of "+cs)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[handSize:], d[:handSize]
}

func (d deck) toString() string {
	stringSlice := []string(d)

	str := strings.Join(stringSlice[:], ",")
	fmt.Println(str)

	return str
}

func toByte(d string) []byte {
	return []byte(d)
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile("./hanem.txt", toByte(d.toString()), 0644)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	stringSlice := strings.Split(string(bs), ",")
	return deck(stringSlice)

}
