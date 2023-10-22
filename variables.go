package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	//card := "Ace of Spades" // Go will infer the type of the variable
	card := NewCard()
	fmt.Println(card)
}

func NewCard() string {
	return "Five of Diamonds"
}
