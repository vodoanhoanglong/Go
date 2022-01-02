package main

import (
	"abstract_factory_example/abstractfactory"
	"fmt"
)

func main() {
	adidasFactory := abstractfactory.GetSportsFactory("adidas")
	adidasShoe := adiasFactory.MakeShoe()
	printShoeDetails(adidasShoe)
	adidasShort := adiasFactory.MakeShoe()
	printShoeDetails(adidasShort)

	nikeFactory := abstractfactory.GetSportsFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	printShoeDetails(nikeShoe)
	nikeShort := nikeFactory.MakeShort()
	printShortDetails(nikeShort)
}

func printShoeDetails(s abstractfactory.IShoe) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %s\n", s.GetSize())
}

func printShortDetails(s abstractfactory.IShort) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %s\n", s.GetSize())
}
