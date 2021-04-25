package main

func main() {
	cards := deck{"One of Diamond", variableTest()}
	cards = append(cards, "Six of Diamonds")

	cardNum := num{1, 2, 3}
	cardNum = append(cardNum, 4)

	cards.print()
	cardNum.printNum()
}
