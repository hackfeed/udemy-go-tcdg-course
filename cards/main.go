package main

func main() {
	cards := newDeckFromFile("save.txt")
	cards.print()
}
