package main

func main() {
	card := deck{"Ace of Diamonds", NewCard()}
	card = append(card, "Six of Spades")

	card.print()

}

func NewCard() string {
	return "Five of Diamonds"
}
