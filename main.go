package main

import "fmt"

var deckSize int

func main() {
	deckLocal := newDeckFromFile("hanem.txt")
	fmt.Println("this is the deck", deckLocal)

}

func newCard() string {
	return "Five of Diamnods"
}
