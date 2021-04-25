package main

import "fmt"

type deck []string
type num []int

func (d deck) print() {
	for i, card := range d{
		fmt.Println(i, card)
	}
}

func (n num) printNum() {
	for i, cardNum := range n{
		fmt.Println(i, cardNum)
	}
}

func variableTest() string {
	return "Six of Diamonds"
}