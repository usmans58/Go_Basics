package main

import "fmt"

func main() {
	card := []string{"Ace of Diamonds", NewCard()}
	card = append(card, "Six of Spades")

	for i, cards := range card {
		fmt.Println(i, cards)
	}

}

func NewCard() string {
	return "Five of Diamonds"
}
