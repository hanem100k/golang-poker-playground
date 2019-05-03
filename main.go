package main

var deckSize int

func main() {
	deckLocal := newDeckFromFile("hanem.txt")
	deckLocal.shuffle()
	deckLocal.print()
}
